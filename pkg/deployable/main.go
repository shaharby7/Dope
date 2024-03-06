package deployable

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"Dope/deployable/constants"
	"Dope/deployable/controllable"
	"Dope/deployable/loggable"
)

type Environments string

const (
	LOCAL Environments = "LOCAL"
)

type DeployableConfig struct {
	ProjectName          string
	RequiredEnvVariables []string
}

type Deployable struct {
	ctx           context.Context
	ENV           Environments
	ProjectName   string
	Logger        loggable.Loggable
	Controllables []controllable.Controllable
	OnError       func(ctx context.Context, err error)
}

func NewDeployable(
	deployableConfig DeployableConfig,
	logger loggable.Loggable,
	onError func(ctx context.Context, err error),
	enfile string, //TODO make it controlled by deployer
) (*Deployable, error) {
	ENV := Environments(os.Getenv("ENV"))
	if "" == ENV {
		ENV = "LOCAL"
	}
	if ENV == LOCAL {
		err := loadEnvfile(deployableConfig.ProjectName, enfile)
		if nil != err {
			fmt.Printf("Warning, could not load local env file: %s\n", err)
		}
	}
	err := verifyEnvVariables(&deployableConfig.RequiredEnvVariables)
	deployable := &Deployable{
		ENV:           ENV,
		ProjectName:   deployableConfig.ProjectName,
		Controllables: make([]controllable.Controllable, 0),
		Logger:        logger,
		OnError:       onError,
	}
	return deployable, err
}

func (deployable *Deployable) Start(parentContext context.Context) {
	var deployableWaitGroup sync.WaitGroup
	ctx := deployable.initiateDeployableContext(parentContext)
	for _, controller := range deployable.Controllables {
		deployableWaitGroup.Add(1)
		err := controller.Start(ctx, &deployableWaitGroup)
		if err != nil {
			deployable.OnError(ctx, err)
		}
	}
	fmt.Printf("Project %s is running", deployable.ProjectName)
	deployableWaitGroup.Wait()
}

func (deployable *Deployable) RegisterControllable(c controllable.Controllable) {
	deployable.Controllables = append(deployable.Controllables, c)
}

func (deployable *Deployable) initiateDeployableContext(parentContext context.Context) context.Context {
	ctx := context.WithValue(parentContext, constants.LOGGER_REF, &deployable.Logger)
	ctx = context.WithValue(ctx, constants.DEPLOYABLE_REF, &deployable)
	deployable.ctx = ctx
	return ctx
}

func verifyEnvVariables(requiredEnvVariables *[]string) error {
	for _, name := range *requiredEnvVariables {
		val := os.Getenv(name)
		if "" == val {
			return errors.New(fmt.Sprintf("required ENV variable %s is not found", name))
		}
	}
	return nil
}

func loadEnvfile(projectName string, envfile string) error {
	// _, b, _, _ := runtime.Caller(0)
	// basePath := filepath.Dir(b)
	// envFilePath := fmt.Sprintf("%s/local/%s.env", basePath, projectName)
	return godotenv.Load(envfile)
}

package cmd

import (
	"fmt"
	"path"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/shaharby7/Dope/pkg/config/entity"
	"github.com/shaharby7/Dope/pkg/utils"

	"github.com/shaharby7/Dope/pkg/config"

	cliUtils "github.com/shaharby7/Dope/pkg/cli/cmd/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmdRoot.AddCommand(cmdCreate)
}

var cmdCreate = &cobra.Command{
	Use:   "create",
	Short: "create new dope object",
	Long:  `use the create command to add new dope objects`,
	RunE: func(cmd *cobra.Command, args []string) error {
		entityManifest, err := getObjectManifest(args)
		if err != nil {
			return utils.FailedBecause("infer dope object type", err)
		}
		createdEntity := &entity.Entity{
			Api:                config.DOPE_CORE_API,
			Type:               entityManifest.Name,
			EntityTypeManifest: entityManifest,
		}
		createdEntity.Name, err = cliUtils.GetInputString("Name", true)
		if err != nil {
			return utils.FailedBecause("getting name", err)
		}
		createdEntity.Description, err = cliUtils.GetInputString("Description", false)
		if err != nil {
			return utils.FailedBecause("getting description", err)
		}
		createdEntity.Binding, err = getObjectBinding(createdEntity)
		if err != nil {
			utils.FailedBecause("getting binding", err)
		}
		values, err := getObjectValues(createdEntity)
		if err != nil {
			utils.FailedBecause("values binding", err)
		}
		createdEntity.Values = values
		entityPath, err := inferObjPath(
			createdEntity,
		)
		if err != nil {
			return err
		}
		fmt.Println(entityPath)
		return entity.WriteEntity(entityPath, createdEntity)
	},
}

func getObjectManifest(args []string) (*entity.EntityTypeManifest, error) {
	if (len(args)) != 1 {
		return nil, fmt.Errorf("could not create command can only be executed with with argument, found %d", len(args))
	}
	name := args[0]
	return entity.GetEntityTypeManifest(config.DOPE_CORE_API, name)
}

func getObjectBinding(e *entity.Entity) (entity.EntityBind, error) {
	manifest := e.EntityTypeManifest
	res := entity.EntityBind{}
	if manifest.BindingSettings == nil {
		return res, nil
	}
	for _, bindType := range manifest.BindingSettings.Must {
		bindName, err := cliUtils.GetInputString(bindType, true)
		if err != nil {
			return nil, err
		}
		res[bindType] = bindName
	}
	for _, bindType := range manifest.BindingSettings.Might {
		bindName, err := cliUtils.GetInputString(bindType, false)
		if err != nil {
			return nil, err
		}
		res[bindType] = bindName
	}
	return res, nil
}

func inferObjPath(
	entity *entity.Entity,
) (string, error) {
	entityManifest := entity.EntityTypeManifest
	templateName := entityManifest.Name
	parsedPathTemplate := template.Must( // TODO: pre-parse template
		template.New(templateName).Funcs(sprig.FuncMap()).Parse(entityManifest.CliOptions.PathTemplate),
	)
	buff, err := utils.ApplyTemplateSafe(parsedPathTemplate, templateName, entity)
	if err != nil {
		return "", utils.FailedBecause(fmt.Sprintf("infer path for entity %s", entity.Name), err)
	}
	return path.Join(
		viper.GetString(string(CONF_VARS_DOPE_PATH)),
		buff.String(),
		fmt.Sprintf("%s.%s.dope.yaml", entity.Name, entity.Type),
	), nil
}

func getObjectValues(_ *entity.Entity) (any, error) {
	return nil, nil
}

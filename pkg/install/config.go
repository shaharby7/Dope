package install

const (
	DEFAULT_NAMESPACE    = "dope"
	DEFAULT_UPGRADE      = false
	DEFAULT_RELEASE_NAME = "dope"
	DEFAULT_ENVIRONMENT  = "local"
)

type config struct {
	ReleaseName string
	Upgrade     bool
	Namespace   string
	Environment string
}

func NewConfig(options ...func(*config)) *config {
	c := &config{
		ReleaseName: DEFAULT_RELEASE_NAME,
		Upgrade:     DEFAULT_UPGRADE,
		Namespace:   DEFAULT_NAMESPACE,
		Environment: DEFAULT_ENVIRONMENT,
	}
	return c
}

func Config_SetReleaseName(name string) func(*config) {
	return func(c *config) {
		c.ReleaseName = name
	}
}

func Config_Upgrade() func(*config) {
	return func(c *config) {
		c.Upgrade = true
	}
}

func Config_SetNamespace(namespace string) func(*config) {
	return func(c *config) {
		c.Namespace = namespace
	}
}

func Config_SetEnvironment(environment string) func(*config) {
	return func(c *config) {
		c.Environment = environment
	}
}
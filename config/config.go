package config

type AppConfig struct {
	DatabaseURL            string
	ApplicationEnvironment string
	LogLevel               string
}

func (ac *AppConfig) Defaults() map[string]string {
	return map[string]string{
		"APPLICATION_ENV": "local",
		"DATABASE_URL":    "",
	}
}

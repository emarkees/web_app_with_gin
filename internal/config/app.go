package config

type AppConfig struct {
	Port string
	Env  string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		Port: getEnv("PORT", "8080"),
		Env: getEnv("APP_ENV", "development"),
	}
}

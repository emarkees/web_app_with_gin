package config

import "os"


type Config struct {
	APP AppConfig
	DB DBConfig
}

func Load() *Config {
	return &Config{
		APP: LoadAppConfig(),
		DB: LoadDatabaseConfig(),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}


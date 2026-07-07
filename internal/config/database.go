package config

import (
	"time"
)

type PostgresConfig struct {
	URL         string
	MaxConns    int
	MinConns    int
	ConnTimeout time.Duration
}

// type MySqlConfig struct {
// 	URL         string
// 	MaxConns    int
// 	MinConns    int
// 	ConnTimeout time.Duration
// }

type DBConfig struct {
	Postgres PostgresConfig
	// Mysql    MySqlConfig
}

func LoadDatabaseConfig() DBConfig {
	return DBConfig{
		Postgres: PostgresConfig{
			URL:         getEnv("DATABASE_URL", "0.0.0.0"),
			MaxConns:    25,
			MinConns:    5,
			ConnTimeout: 5 * time.Second,
		},

		// MySQL database configuration
		// Mysql: MySqlConfig{
		// 	URL:         getEnv("DATABASE_URL", "1.0.0.1"),
		// 	MaxConns:    25,
		// 	MinConns:    5,
		// 	ConnTimeout: 5 * time.Second,
		// },
	}
}

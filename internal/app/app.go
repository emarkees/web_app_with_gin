package app

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// This check the postgres to see if a database exist, if it exist it proceed with migrate
// if not it autocreate DB
// func EnsureDBExist(dbUser, dbPass, dbHost, dbPort, dbName, sslMode string) error {
// 	mainConns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
// 		dbUser, dbPass, dbHost, dbPort, "postgres", sslMode,
// 	)

// 	pool, err := pgxpool.New(context.Background(), mainConns)
// 	if err != nil {
// 		return err
// 	}

// 	defer pool.Close()

// 	var exists bool

// 	query := "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)"
// 	err = pool.QueryRow(context.Background(), query, dbName).Scan(&exists)
// 	if err != nil {
// 		return fmt.Errorf("failed to check if database exists: %v", err)
// 	}

// 	if !exists {
// 		_, err := pool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s", dbName))
// 		if err != nil {
// 			return fmt.Errorf("failed to create database: %v", err)
// 		}
// 		fmt.Printf("Database '%s', created successfully.\n", dbName)

// 	} else {
// 		fmt.Printf("Database '%s' already exist.\n", dbName)
// 	}

// 	return nil
// }

// ConnectDB opens a PostgreSQL connection pool using the provided DSN.
// The caller is responsible for calling pool.Close() when done.
func ConnectDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Verify the connection is alive
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}

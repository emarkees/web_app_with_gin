package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunMigrations(pool *pgxpool.Pool, migrationsDir string) error {
	ctx := context.Background()

	// 1. Create the migrations tracking table
	createTableQuery := `CREATE TABLE IF NOT EXISTS schema_migrations (
		filename TEXT PRIMARY KEY,
		applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`

	_, err := pool.Exec(ctx, createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// 2. Read the directory
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		var applied bool

		// FIX: Corrected table name to match creation ('schema_migrations' plural)
		checkQuery := `SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE filename = $1)`

		err := pool.QueryRow(ctx, checkQuery, file.Name()).Scan(&applied)
		if err != nil {
			return fmt.Errorf("failed checking migration status for %s: %w", file.Name(), err)
		}

		if applied {
			log.Printf("Skipping (already applied): %s", file.Name())
			continue
		}

		// FIX: Use os.ReadFile to read file content, not os.ReadDir
		content, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file.Name(), err)
		}

		// Split the file content, take the first part (the Up section)
		parts := strings.Split(string(content), "-- +migrate Down")
		upSQL := strings.Replace(parts[0], "-- +migrate Up", "", 1)
		upSQL = strings.TrimSpace(upSQL)

		// MISSING LOGIC ADDED: Execute the migration and record it inside a transaction
		log.Printf("Applying migration: %s", file.Name())
		
		tx, err := pool.Begin(ctx)
		if err != nil {
			return fmt.Errorf("failed to begin transaction for %s: %w", file.Name(), err)
		}

		// Execute the migration logic
		if _, err := tx.Exec(ctx, upSQL); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
		}

		// Record the migration in the tracking table
		insertQuery := `INSERT INTO schema_migrations (filename) VALUES ($1)`
		if _, err := tx.Exec(ctx, insertQuery, file.Name()); err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to record migration %s: %w", file.Name(), err)
		}

		// Commit the transaction
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("failed to commit transaction for %s: %w", file.Name(), err)
		}
		
		log.Printf("Successfully applied: %s", file.Name())
	}

	return nil
}
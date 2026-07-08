package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ Error: Please provide a migration name.")
		fmt.Println("Usage: make make:migration <migration_name>")
		os.Exit(1)
	}

	migrationName := strings.ToLower(os.Args[1])
	migrationsDir := filepath.Join("db", "migrations")

	// 1. Ensure the directory exists
	if err := os.MkdirAll(migrationsDir, os.ModePerm); err != nil {
		fmt.Printf("❌ Error creating migrations directory: %v\n", err)
		os.Exit(1)
	}

	// 2. Check if a migration with this exact name suffix already exists
	files, err := os.ReadDir(migrationsDir)
	if err == nil { // Only check if we can read the directory successfully
		expectedSuffix := "_" + migrationName + ".sql"
		for _, file := range files {
			if strings.HasSuffix(file.Name(), expectedSuffix) {
				fmt.Printf("⚠️  Error: A migration named '%s' already exists!\n", migrationName)
				fmt.Printf("   See: %s\n", filepath.Join(migrationsDir, file.Name()))
				os.Exit(1) // Stop execution immediately
			}
		}
	}

	// 3. If no duplicate is found, proceed to create the file
	version := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s.sql", version, migrationName)
	filePath := filepath.Join(migrationsDir, fileName)

	tableName := extractTableName(migrationName)
	
	boilerplate := fmt.Sprintf(`-- +migrate Up
CREATE TABLE IF NOT EXISTS %s (
    id INT PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- +migrate Down
DROP TABLE IF EXISTS %s;
`, tableName, tableName)

	if err := os.WriteFile(filePath, []byte(boilerplate), 0644); err != nil {
		fmt.Printf("❌ Error writing migration file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("🚀 Unified migration template scaffolded successfully!\n")
	fmt.Printf("  📄 %s\n", filePath)
}

func extractTableName(name string) string {
	parts := strings.Split(name, "_")
	if len(parts) >= 2 && parts[0] == "create" {
		if parts[len(parts)-1] == "table" {
			return strings.Join(parts[1:len(parts)-1], "_")
		}
		return strings.Join(parts[1:], "_")
	}
	return "my_table"
}
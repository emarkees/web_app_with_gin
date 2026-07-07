package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/emarkees/internal/app"
	"github.com/emarkees/internal/config"
	"github.com/emarkees/internal/models"
	"github.com/emarkees/internal/route"
	"github.com/joho/godotenv"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	dsn := flag.String("dsn", "", "PostgreSQL connection string")
	flag.Parse()

	// 1. Log cleanly to standard streams (Stdout/Stderr)
	// No files are opened here inside Go!
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Loading the .env files
	if err := godotenv.Load(); err != nil {
		errorLog.Printf("Warning: .env file not found, pulling from system enviroment variables: %v", err)
	}

	// 2. Connect to PostgreSQL
	if *dsn == "" {
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")
		sslMode := os.Getenv("DB_SSLMODE")

		defaultDSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser, dbPass, dbHost, dbPort, dbName, sslMode,
		)

		dsn = &defaultDSN
	}

	pool, err := app.ConnectDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pool.Close()
	infoLog.Println("Database connection pool established")

	// Initialize a models.SnippetModel instance
	snippetsModel := &models.SnippetModel{DB: pool}

	// 3. Build the dependency container
	ctrlx := config.NewContainer(infoLog, errorLog, pool, snippetsModel)
	r := route.SetRoute(ctrlx)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  r,
	}

	infoLog.Printf("Server is running on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}

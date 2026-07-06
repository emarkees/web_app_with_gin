package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/emarkees/internal/app"
	"github.com/emarkees/internal/config"
	"github.com/emarkees/internal/route"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	dsn := flag.String("dsn", "", "PostgreSQL connection string")
	flag.Parse()

	// 1. Log cleanly to standard streams (Stdout/Stderr)
	// No files are opened here inside Go!
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// 2. Connect to PostgreSQL
	if *dsn == "" {
		// Default DSN for local development — override with -dsn flag
		defaultDSN := "postgres://postgres:password@localhost:5432/snippetbox?sslmode=disable"
		dsn = &defaultDSN
	}

	pool, err := app.ConnectDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pool.Close()
	infoLog.Println("Database connection pool established")

	// 3. Build the dependency container
	ctrlx := config.NewContainer(infoLog, errorLog, pool)
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
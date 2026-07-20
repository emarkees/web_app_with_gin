package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/emarkees/internal/app"
	"github.com/emarkees/internal/config"
	"github.com/emarkees/internal/models"
	"github.com/emarkees/internal/route"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dns := flag.String("dns", "postgres://web:web@localhost:5432/snippetbox?sslmode=disable", "POSTGRES source name")
	flag.Parse()

	// log both informations and errors
	f, err := os.OpenFile("storage/logs/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err := app.ConnectDB(*dns)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer pool.Close()

	snippetModel := &models.SnippetModel{DB: pool}
	// Build a dependency container
	ctrlx := config.NewContainer(infoLog, errorLog, pool, snippetModel)

	r := route.SetRoute(ctrlx)

	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  r,
	}

	infoLog.Printf("Server is running on port %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}

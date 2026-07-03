package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/emarkees/config"
	// "github.com/emarkees/internal/controller"
	"github.com/emarkees/internal/route"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	flag.Parse()

	// 1. Log cleanly to standard streams (Stdout/Stderr)
	// No files are opened here inside Go!
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	ctrlx := config.NewContainer()(infoLog, errorLog)
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
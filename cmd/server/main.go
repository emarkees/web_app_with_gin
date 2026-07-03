package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/emarkees/internal/controller"
	"github.com/emarkees/internal/route"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	flag.Parse()

	// Create or open a log file
	f, err := os.OpenFile("/temp/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening info log file: %v", err)
	}

	defer f.Close()

	// Use log.New() to create a logger for writing information messages.
	infoLog := os.New(f, "INFO\t", log.Ldate|log.Ltime)

	errorLog := os.New(f, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	ctrlx := controller.NewContainer()

	r := route.SetRoute(ctrlx)


	//Initialise a serverHTTP server struct
	srv := &http.Server{
		Addr : *addr,
		ErrorLog: errorLog,
		Handler: r,
	}

	infoLog.Printf("Server is running on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}

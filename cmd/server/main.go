package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/emarkees/internal/controller"
	"github.com/emarkees/internal/route"
)

func main() {
	Addr := flag.String("port", ":4000", "Usage: go run ./cmd/server")
	flag.Parse()

	ctrlx := controller.NewContainer()

	r := route.SetRoute(ctrlx)


	
	log.Println("Server is running on port", *Addr)
	if err := http.ListenAndServe(*Addr, r); err != nil {
		log.Fatal(err)
	}
}
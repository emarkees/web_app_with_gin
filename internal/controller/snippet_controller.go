package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
)

type Container struct {
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Home(w http.ResponseWriter, r *http.Request) {

	tsx, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tsx.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func (c *Container) Store(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my create page!\n"))
}


func (c *Container) Show(w http.ResponseWriter, r *http.Request) {
	idx := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idx)
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "Displaying the snippet with id %d", id)
}

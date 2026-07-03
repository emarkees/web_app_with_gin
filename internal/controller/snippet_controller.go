package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/emarkees/config"
	"github.com/go-chi/chi/v5"
)

type RouterContext struct {
	App *config.Container
}

func (c *RouterContext) Home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	tsx, err := template.ParseFiles(files...)
	if err != nil {
		c.ctx.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tsx.ExecuteTemplate(w, "base", nil)
	if err != nil {
		c.ctx.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func (c *RouterContext) Store(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my create page!\n"))

}

func (c *RouterContext) Show(w http.ResponseWriter, r *http.Request) {
	idx := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idx)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying the snippet with id %d", id)
}

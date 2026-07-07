package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/emarkees/internal/config"
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
		c.App.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = tsx.ExecuteTemplate(w, "base", nil)
	if err != nil {
		c.App.ErrorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func (c *RouterContext) Store(w http.ResponseWriter, r *http.Request) {
	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7
	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := c.App.snippets.Insert(title, content, expires)
	if err != nil {
		c.App.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

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

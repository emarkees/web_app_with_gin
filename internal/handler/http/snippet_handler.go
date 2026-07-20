package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	// "text/template"

	"github.com/emarkees/internal/config"
	"github.com/emarkees/internal/models"
	"github.com/go-chi/chi/v5"
)

type RouterContext struct {
	App *config.Container
}

func (c *RouterContext) Home(w http.ResponseWriter, r *http.Request) {

	snippets, err := c.App.Snippets.Latest(context.Background())
	if err != nil {
		c.ServerError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v", snippet)
	}

	// files := []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// }

	// // template.ParseFiles is used to 
	// tsx, err := template.ParseFiles(files...)
	// if err != nil {
	// 	c.ServerError(w, err)
	// 	return
	// }

	// err = tsx.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	c.ServerError(w, err)
	// 	return
	// }

}

func (c *RouterContext) Store(w http.ResponseWriter, r *http.Request) {
	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := c.App.Snippets.Insert(r.Context(), title, content, expires)
	if err != nil {
		c.ServerError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

}

func (c *RouterContext) Show(w http.ResponseWriter, r *http.Request) {
	idx := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idx)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := c.App.Snippets.Get(context.Background(), id)
	if err != nil {
		if errors.Is(err, models.ErrorRecord) {
			c.NotFound(w)
		} else {
			c.ServerError(w, err)
		}
	}

	fmt.Fprintf(w, "%+v", snippet)
}

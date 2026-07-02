package route

import (
	"github.com/emarkees/internal/controller"
	"github.com/go-chi/chi/v5"
)

func SetRoute(ctrlx *controller.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", ctrlx.Home)
	r.Post("/create", ctrlx.Store)
	r.Get("/snippet/{id}", ctrlx.Show)

	return r
}

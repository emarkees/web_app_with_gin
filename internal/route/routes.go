package route

import (
	"github.com/emarkees/internal/controller"
	"github.com/go-chi/chi/v5"
)

func SetRoute(ctrlx *controller.Container) *chi.Mux {
	r := chi.NewMux()

	r.Get("/", ctrlx.Home)

	return r
}

package route

import (
	"net/http"
	"path/filepath"

	"github.com/emarkees/internal/controller"
	"github.com/go-chi/chi/v5"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, err
			}

			return nil, err
		}
	}

	return f, nil
}

func SetRoute(ctrlx *controller.Container) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", ctrlx.Home)
	r.Post("/create", ctrlx.Store)
	r.Get("/snippet/{id}", ctrlx.Show)

	fs := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	r.Handle("/static", http.NotFoundHandler())
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	return r
}

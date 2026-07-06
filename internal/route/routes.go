package route

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/emarkees/internal/config"
	handler "github.com/emarkees/internal/handler/http"
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
		f.Close()
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {

			f.Close()

			return nil, os.ErrPermission
		}
	}

	return f, nil
}

func SetRoute(cfg *config.Container) *chi.Mux {
	r := chi.NewRouter()

	// Initialize the shared context for all your handlers
	ctx := &handler.RouterContext{App: cfg}

	r.Get("/", ctx.Home)
	r.Post("/create", ctx.Store)
	r.Get("/snippet/{id}", ctx.Show)

	fs := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	r.Handle("/static", http.NotFoundHandler())
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	return r
}

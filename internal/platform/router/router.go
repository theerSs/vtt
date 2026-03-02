package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ApiHandler interface {
	Register(chi.Router)
}

type Deps struct {
	APIs map[string]ApiHandler
}

func InitRouter(d Deps) *chi.Mux{
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	
	r.Route("/api", func(r chi.Router) {
		for path, h := range d.APIs {
			r.Route("/"+path, func(r chi.Router) {
				h.Register(r)
			})
		}
	})

	return r
}
package router

import (
	"github.com/go-chi/chi/v5"
)

type AppModule interface {
	RegisterRoutes(chi.Router)
}

type Deps struct {
	Modules map[string]AppModule
}

func InitRouter(d Deps) *chi.Mux {
	r := chi.NewRouter()
	setupMiddlewares(r)

	r.Route("/api", func(r chi.Router) {
		for path, m := range d.Modules {
			r.Route("/"+path, func(r chi.Router) {
				m.RegisterRoutes(r)
			})
		}
	})

	return r
}

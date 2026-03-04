package rooms

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/theerSs/vtt/internal/platform/router"
)

type module struct {
	handler *handler
}

func NewModule(db *pgxpool.Pool) router.AppModule {
	repository := newRepository(db)
	service := newService(repository)
	handler := newHandler(service)

	return &module{
		handler: handler,
	}
}

func (m *module) RegisterRoutes(r chi.Router) {
	r.Get("/", m.handler.GetList)
}

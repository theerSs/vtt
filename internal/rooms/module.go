package rooms

import (
	"github.com/go-chi/chi/v5"
	"github.com/theerSs/vtt/internal/platform/router"
)

type module struct {
	handler *handler
}

func NewModule() router.AppModule {
	repository := newRepository()
	service := newService(repository)
	handler := newHandler(service)

	return &module{
		handler: handler,
	}
}

func (m *module) RegisterRoutes(r chi.Router) {
	r.Get("/", m.handler.GetList)
}

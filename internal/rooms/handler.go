package rooms

import (
	"fmt"
	"net/http"
)

type handler struct {
	service *service
}

func newHandler(s *service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLED")
}

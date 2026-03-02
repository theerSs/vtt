package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}


func (h *Handler) Register(r chi.Router) {
	r.Get("/login", h.Login)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("login"))
}
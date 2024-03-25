package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterest/go-basic-ecom/types"
)

type Handler struct {
	repository types.UserRepository
}

func NewHandler(repository types.UserRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.loginController).Methods(http.MethodPost)
	router.HandleFunc("/register", h.registrationController).Methods(http.MethodPost)
}

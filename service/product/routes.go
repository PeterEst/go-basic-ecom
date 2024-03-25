package product

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterest/go-basic-ecom/types"
)

type Handler struct {
	repository types.ProductRepository
}

func NewHandler(repository types.ProductRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
	router.HandleFunc("/products/{id}", h.handleGetProduct).Methods(http.MethodGet)
}

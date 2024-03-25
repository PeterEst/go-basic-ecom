package cart

import (
	"github.com/gorilla/mux"
	"github.com/peterest/go-basic-ecom/types"
)

type Handler struct {
	productRepository types.ProductRepository
	orderRepository   types.OrderRepository
	userRepository    types.UserRepository
}

func NewHandler(
	productRepository types.ProductRepository,
	orderRepository types.OrderRepository,
	userRepository types.UserRepository,
) *Handler {
	return &Handler{
		productRepository: productRepository,
		orderRepository:   orderRepository,
		userRepository:    userRepository,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.checkoutController).Methods("POST")
}

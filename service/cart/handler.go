package cart

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/peterest/go-basic-ecom/middleware"
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
	router.HandleFunc("/cart/checkout", middleware.WithJWTAuth(h.checkoutController, h.userRepository)).Methods(http.MethodPost)
}

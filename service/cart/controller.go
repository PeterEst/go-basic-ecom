package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/peterest/go-basic-ecom/middleware"
	"github.com/peterest/go-basic-ecom/types"
	"github.com/peterest/go-basic-ecom/utils"
)

func (h *Handler) checkoutController(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserIDFromContext(r.Context())

	var cart types.CheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validator.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid payload: %s", errors))
		return
	}

	productIds, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.productRepository.GetProductsByID(productIds)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSuccessfulAPIResponse(w, map[string]interface{}{
		"orderId":    orderID,
		"totalPrice": totalPrice,
	}, "success", http.StatusOK)
}

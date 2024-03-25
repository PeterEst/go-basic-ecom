package cart

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/peterest/go-basic-ecom/types"
	"github.com/peterest/go-basic-ecom/utils"
)

// TODO: authentication
func (h *Handler) checkoutController(w http.ResponseWriter, r *http.Request) {
	userID := 4

	var cart types.CheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	productIds, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	products, err := h.productRepository.GetProductsByID(productIds)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"orderId":    orderID,
		"totalPrice": totalPrice,
	})
}

package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/peterest/go-basic-ecom/types"
	"github.com/peterest/go-basic-ecom/utils"
)

func (h *Handler) getProductsController(w http.ResponseWriter, r *http.Request) {
	products, err := h.repository.GetProducts()

	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccessfulAPIResponse(w, products, "success", http.StatusOK)
}

func (h *Handler) createProductController(w http.ResponseWriter, r *http.Request) {
	var product types.CreateProductPayload
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validator.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid payload: %s", errors))
		return
	}

	createdProduct, err := h.repository.CreateProduct(product)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccessfulAPIResponse(w, createdProduct, "success", http.StatusCreated)
}

func (h *Handler) getProductController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, "missing product id")
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, "invalid product id")
		return
	}

	product, err := h.repository.GetProductByID(productID)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSuccessfulAPIResponse(w, product, "success", http.StatusOK)
}

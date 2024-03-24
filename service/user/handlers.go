package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/peterest/go-basic-ecom/types"
	"github.com/peterest/go-basic-ecom/utils"
)

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get JSON Payload
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Validate Payload
	if err := utils.Validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	// Check if user exists
	_, err := h.repository.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// Basic Hash Password
	hashedPassword, err := utils.Hash(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Create user
	err = h.repository.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

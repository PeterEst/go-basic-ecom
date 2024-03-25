package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/peterest/go-basic-ecom/types"
	"github.com/peterest/go-basic-ecom/utils"
)

func (h *Handler) loginController(w http.ResponseWriter, r *http.Request) {
	var credentials types.LoginUserPayload
	if err := utils.ParseJSON(r, &credentials); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator.Struct(credentials); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	user, err := h.repository.GetUserByEmail(credentials.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid credentials"))
		return
	}

	if !utils.CompareHash(user.Password, credentials.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid credentials"))
		return
	}

	tokenClaims := map[string]interface{}{"userID": user.ID}
	var expiration time.Duration = time.Second * time.Duration(60*60*24*7) // 1 week

	token, err := utils.GenerateJWT(tokenClaims, expiration)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) registrationController(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	_, err := h.repository.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	hashedPassword, err := utils.Hash(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

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

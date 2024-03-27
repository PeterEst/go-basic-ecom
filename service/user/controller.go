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
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, "failed")
		return
	}

	if err := utils.Validator.Struct(credentials); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid payload: %s", errors))
		return
	}

	user, err := h.repository.GetUserByEmail(credentials.Email)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, "invalid credentials")
		return
	}

	if !utils.CompareHash(user.Password, credentials.Password) {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, "invalid credentials")
		return
	}

	tokenClaims := map[string]interface{}{"userID": user.ID}
	var expiration time.Duration = time.Second * time.Duration(60*60*24*7) // 1 week

	token, err := utils.GenerateJWT(tokenClaims, expiration)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusInternalServerError, "failed")
		return
	}

	utils.HandleSuccessfulAPIResponse(w, map[string]string{"token": token}, "success", http.StatusOK)
}

func (h *Handler) registrationController(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, "failed")
		return
	}

	if err := utils.Validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.HandleFailedAPIResponse(w, http.StatusBadRequest, fmt.Sprintf("invalid payload: %s", errors))
		return
	}

	_, err := h.repository.GetUserByEmail(user.Email)
	if err == nil {
		utils.HandleFailedAPIResponse(w, http.StatusConflict, fmt.Sprintf("user with email %s already exists", user.Email))
		return
	}

	hashedPassword, err := utils.Hash(user.Password)
	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.repository.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.HandleFailedAPIResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSuccessfulAPIResponse(w, nil, "success", http.StatusCreated)
}

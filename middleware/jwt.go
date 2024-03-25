package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/peterest/go-basic-ecom/types"
	"github.com/peterest/go-basic-ecom/utils"
)

type CtxKey string

const UserKey CtxKey = "userID"

func WithJWTAuth(handlerFunc http.HandlerFunc, userRepository types.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetJWTFromRequest(r)

		token, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID, ok := claims["userID"].(float64)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := userRepository.GetUserByID(int(userID))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, user.ID)

		handlerFunc(w, r.WithContext(ctx))
	}
}

func GetUserIDFromContext(ctx context.Context) int {
	userID := ctx.Value(UserKey)

	if userID == nil {
		return -1
	}

	return userID.(int)
}

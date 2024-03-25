package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/peterest/go-basic-ecom/config"
)

var (
	secretKey = config.Env.JwtSecret
)

func GenerateJWT(claims map[string]interface{}, expiration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims["expiredAt"] = time.Now().Add(expiration).Unix()
	token.Claims = jwt.MapClaims(claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetJWTFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if token != "" {
		// removing "Bearer " prefix
		token = token[7:]

		return token
	}

	return ""
}

func ValidateJWT(token string) (*jwt.Token, error) {
	jwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwt.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return jwt, nil
}

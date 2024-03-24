package utils

import (
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

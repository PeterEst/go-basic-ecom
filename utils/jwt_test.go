package utils

import "testing"

func TestGenerateJWT(t *testing.T) {
	claims := map[string]interface{}{
		"email": "test@gmail.com",
	}

	token, err := GenerateJWT(claims, 1)
	if err != nil {
		t.Error(err)
	}

	if token == "" {
		t.Error("Token should not be empty")
	}
}

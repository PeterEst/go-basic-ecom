package utils

import "golang.org/x/crypto/bcrypt"

func Hash(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHash(hash, str string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(str)) == nil
}

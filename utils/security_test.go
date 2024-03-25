package utils

import "testing"

func TestHash(t *testing.T) {
	password := "password"

	hash, err := Hash(password)
	if err != nil {
		t.Error(err)
	}

	if hash == "" {
		t.Error("Hash should not be empty")
	}

	if hash == password {
		t.Error("Hash should not be same as password")
	}
}

func TestCompareHash(t *testing.T) {
	password := "password"

	hash, err := Hash(password)
	if err != nil {
		t.Error(err)
	}

	if !CompareHash(hash, password) {
		t.Error("Hash should be same as password")
	}
}

package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(12)
	hashedPassword, err := GenerateFromPassword(password)
	if err != nil {
		t.Error(err)
	}
	err = VerifyPassword(password, hashedPassword)
	if err != nil {
		t.Error(err)
	}

	wrongPassword := RandomString(12)
	err = VerifyPassword(wrongPassword, hashedPassword)
	if err == nil {
		t.Error(err)
	}
}

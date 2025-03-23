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
	isMatch, err := VerifyPassword(password, hashedPassword)
	if err != nil {
		t.Error(err)
	}
	if !isMatch {
		t.Error("hash function's incorrect")
	}

	wrongPassword := RandomString(12)
	isMatch, err = VerifyPassword(wrongPassword, hashedPassword)
	if err != nil {
		t.Error(err)
	}
	if isMatch {
		t.Error("hash function's incorrect")
	}
}

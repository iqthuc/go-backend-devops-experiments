package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

// sử dụng thuật toán hash argon2 trong ứng dụng thực tế

const saltSize = 32

func generateSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func GenerateFromPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(append(salt, []byte(password)...))
	return fmt.Sprintf("%s$%s", base64.StdEncoding.EncodeToString(salt), base64.StdEncoding.EncodeToString(hash[:])), nil
}

func VerifyPassword(inputPassword string, storeHash string) (bool, error) {
	parts := strings.Split(storeHash, "$")
	if len(parts) != 2 {
		return false, errors.New("Invalid hash format")
	}
	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err
	}
	expectedHash, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, nil
	}
	inputHash := sha256.Sum256(append(salt, []byte(inputPassword)...))
	return subtle.ConstantTimeCompare(inputHash[:], expectedHash) == 1, nil
}

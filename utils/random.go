package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) string {
	str := make([]byte, 32)
	_, err := rand.Read(str)
	if err != nil {
		log.Fatal("Failed to generate random string:", err)
	}
	return string(str)
}

func GenerateSecretKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("Failed to generate secret key:", err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

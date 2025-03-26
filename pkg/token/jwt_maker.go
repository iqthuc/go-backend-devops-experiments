package token

import (
	"fmt"
	"log"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) Maker {
	return &JWTMaker{
		secretKey: secretKey,
	}
}

func (j *JWTMaker) CreateToken(payload Payload) (string, error) {
	token, err := EncodeJWT(payload, j.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode JWT: %w", err)
	}
	return token, nil
}

func (j *JWTMaker) VerifyToken(token string) (Payload, error) {
	claims, err := DecodeJWT(token, j.secretKey)
	if err != nil {
		log.Println("failed to decode JWT:", err)
		return nil, Errors.InvalidToken
	}
	if err := claims.IsValid(); err != nil {
		return nil, err
	}
	return claims, nil
}

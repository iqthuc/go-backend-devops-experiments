package token

import (
	"fmt"
	"log"
	"time"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) Maker {
	return &JWTMaker{
		secretKey: secretKey,
	}
}

func (j *JWTMaker) CreateToken(userId int64, duration time.Duration) (string, error) {
	claims := Payload{
		UserId: userId,
		Exp:    time.Now().Add(duration),
	}

	token, err := EncodeJWT(claims, j.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to encode JWT: %w", err)
	}
	return token, nil
}

func (j *JWTMaker) VerifyToken(token string) (*Payload, error) {
	claims, err := DecodeJWT(token, j.secretKey)
	if err != nil {
		log.Println("failed to decode JWT:", err)
		return nil, Errors.InvalidToken
	}
	if time.Now().After(claims.Exp) {
		return nil, Errors.ExpiredToken
	}
	return claims, nil
}

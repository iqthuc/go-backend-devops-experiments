package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/iqthuc/sport-shop/config"
)

var JWTErrors = struct {
	ErrInvalidToken error
	ErrExpiredToken error
}{
	ErrInvalidToken: errors.New("invalid token"),
	ErrExpiredToken: errors.New("token expired"),
}
var secretKey = os.Getenv(config.Env.SecretKey)

type JWTClaims struct {
	UserId int64     `json:"user_id"`
	Exp    time.Time `json:"exp"`
}

func generateHMAC(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func EncodeJWT(claims JWTClaims) (string, error) {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))

	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)

	signature := generateHMAC(header+"."+payload, secretKey)
	return header + "." + payload + "." + signature, nil
}

func DecodeJWT(token string) (*JWTClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, JWTErrors.ErrInvalidToken
	}

	headerAndPayload := parts[0] + "." + parts[1]
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, JWTErrors.ErrInvalidToken
	}

	// tạo lại chữ kí từ header + payload + secretKey
	// và kiểm tra xem có khớp không
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(headerAndPayload))
	expectedSignature := h.Sum(nil)

	if !hmac.Equal(signature, expectedSignature) {
		return nil, JWTErrors.ErrInvalidToken
	}

	// ok, decode payload

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, JWTErrors.ErrInvalidToken
	}

	var claims JWTClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, JWTErrors.ErrInvalidToken
	}

	if time.Now().After(claims.Exp) {
		return nil, JWTErrors.ErrExpiredToken
	}
	return &claims, nil

}

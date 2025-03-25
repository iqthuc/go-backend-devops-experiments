package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

func generateHMAC(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

func EncodeJWT(claims Payload, secretKey string) (string, error) {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))

	payloadBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	payload := base64.RawURLEncoding.EncodeToString(payloadBytes)

	signature := generateHMAC(header+"."+payload, secretKey)
	return header + "." + payload + "." + signature, nil
}

func DecodeJWT(token string, secretKey string) (*Payload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, Errors.InvalidToken
	}

	headerAndPayload := parts[0] + "." + parts[1]
	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, Errors.InvalidToken
	}

	// tạo lại chữ kí từ header + payload + secretKey
	// và kiểm tra xem có khớp không
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(headerAndPayload))
	expectedSignature := h.Sum(nil)

	if !hmac.Equal(signature, expectedSignature) {
		return nil, Errors.InvalidToken
	}

	// ok, decode payload

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, Errors.InvalidToken
	}

	var claims Payload
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, Errors.InvalidToken
	}

	if time.Now().After(claims.Exp) {
		return nil, Errors.ExpiredToken
	}
	return &claims, nil
}

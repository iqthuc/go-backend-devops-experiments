package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
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

func DecodeJWT(token string, secretKey string) (Payload, error) {
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

	// check token type
	var raw map[string]any
	if err := json.Unmarshal(payload, &raw); err != nil {
		return nil, Errors.InvalidToken
	}
	typeStr, ok := raw["token_type"].(float64)

	if !ok {
		return nil, Errors.InvalidToken
	}

	var claims Payload

	switch int(typeStr) {
	case int(Access):
		access := &AccessPayload{}
		if err := json.Unmarshal(payload, access); err != nil {
			return nil, Errors.InvalidToken
		}
		claims = access
	case int(Refresh):
		refresh := &RefreshPayload{}
		if err := json.Unmarshal(payload, refresh); err != nil {
			return nil, Errors.InvalidToken
		}
		claims = refresh
	default:

		return nil, Errors.InvalidToken
	}
	return claims, nil
}

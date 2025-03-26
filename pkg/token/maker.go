package token

import (
	"errors"
)

var Errors = struct {
	InvalidToken error
	ExpiredToken error
}{
	InvalidToken: errors.New("invalid token"),
	ExpiredToken: errors.New("token expired"),
}

type Maker interface {
	CreateToken(payload Payload) (string, error)
	VerifyToken(token string) (Payload, error)
}

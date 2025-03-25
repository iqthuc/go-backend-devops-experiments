package token

import (
	"errors"
	"time"
)

var Errors = struct {
	InvalidToken error
	ExpiredToken error
}{
	InvalidToken: errors.New("invalid token"),
	ExpiredToken: errors.New("token expired"),
}

type Payload struct {
	UserId int64     `json:"user_id"`
	Exp    time.Time `json:"exp"`
}

type Maker interface {
	CreateToken(userId int64, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

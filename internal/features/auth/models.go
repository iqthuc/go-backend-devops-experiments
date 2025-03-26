package auth

import (
	"errors"
	"time"
)

type User struct {
	id           int64
	username     string
	email        string
	passwordHash string
	fullName     string
	phoneNumber  string
	isActive     bool
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
}

type RegisterRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type LoginStatus int

const (
	LoginStatusSuccess LoginStatus = iota
	LoginStatusUserNotFound
	LoginStatusInvalidPassword
	LoginStatusSystemError
)

type loginResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

var Err = struct {
	ErrUserNotFound error
}{
	ErrUserNotFound: errors.New("user not found"),
}

type getUserInfoResult struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type refreshTokenRequest struct {
	UserID     int64     `json:"user_id"`
	Token      string    `json:"token"`
	Expires_at time.Time `json:"expires_at"`
}

type getRefreshTokenResult struct {
	UserID     int64
	Token      string
	Expires_at time.Time
}

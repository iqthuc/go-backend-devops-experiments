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

type loginUserReponse struct {
	AccessToken string `json:"access_token"`
}

type userInfoResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

var Err = struct {
	ErrUserNotFound error
}{
	ErrUserNotFound: errors.New("user not found"),
}

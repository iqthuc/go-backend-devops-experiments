package auth

import "time"

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

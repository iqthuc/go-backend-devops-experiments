package auth

import (
	"context"
	"errors"

	"github.com/iqthuc/sport-shop/utils"
)

type UserCase interface {
	RegisterUser(ctx context.Context, user *User) error
}

type userCase struct {
	repo Repository
}

func NewUserCase(repo Repository) UserCase {
	return &userCase{
		repo: repo,
	}
}

func (u *userCase) RegisterUser(ctx context.Context, user *User) error {
	existingUser, err := u.repo.GetUserByEmail(ctx, user.email)
	if (existingUser != User{}) {
		return errors.New("Email already exists")
	}
	hashedPassword, err := utils.GenerateFromPassword(user.passwordHash)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	user.passwordHash = string(hashedPassword)
	return u.repo.CreateUser(ctx, user)
}

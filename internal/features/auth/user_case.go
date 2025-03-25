package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iqthuc/sport-shop/pkg/token"
	"github.com/iqthuc/sport-shop/utils"
)

type UserCase interface {
	RegisterUser(ctx context.Context, user *User) error
	Login(ctx context.Context, identifier LoginRequest) (*loginUserReponse, LoginStatus, error)
	// để tạm để test auth, chuyển về feature User sau
	GetUserInfo(ctx context.Context, id int64) (*userInfoResponse, error)
}

type userCase struct {
	repo  Repository
	maker token.Maker
}

func NewUserCase(repo Repository, maker token.Maker) UserCase {
	return &userCase{
		repo:  repo,
		maker: maker,
	}
}

func (u *userCase) RegisterUser(ctx context.Context, user *User) error {
	existingUser, err := u.repo.GetUserByUsernamOrEmail(ctx, user.username, user.email)
	if existingUser != nil {
		return errors.New("Email or username already in use")
	}
	hashedPassword, err := utils.GenerateFromPassword(user.passwordHash)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	user.passwordHash = string(hashedPassword)
	return u.repo.CreateUser(ctx, user)
}

func (u *userCase) Login(ctx context.Context, rq LoginRequest) (*loginUserReponse, LoginStatus, error) {
	user, err := u.repo.GetUserByUsernamOrEmail(ctx, rq.Identifier, rq.Identifier)
	if err != nil {
		return nil, LoginStatusSystemError, nil
	}
	if user == nil {
		return nil, LoginStatusUserNotFound, nil
	}

	err = utils.VerifyPassword(rq.Password, user.passwordHash)
	if err != nil {
		return nil, LoginStatusInvalidPassword, nil
	}

	token, err := u.maker.CreateToken(user.id, time.Hour)

	response := &loginUserReponse{
		AccessToken: token,
	}
	return response, LoginStatusSuccess, nil
}

func (u *userCase) GetUserInfo(ctx context.Context, id int64) (*userInfoResponse, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w ", err)
	}
	if user == nil {
		return nil, Err.ErrUserNotFound
	}
	info := &userInfoResponse{
		Username:    user.username,
		Email:       user.email,
		FullName:    user.fullName,
		PhoneNumber: user.phoneNumber,
	}
	return info, nil
}

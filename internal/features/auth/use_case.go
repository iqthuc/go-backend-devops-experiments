package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iqthuc/sport-shop/pkg/token"
	"github.com/iqthuc/sport-shop/utils"
)

type UseCase interface {
	RegisterUser(ctx context.Context, user *User) error
	Login(ctx context.Context, identifier LoginRequest) (*loginResult, LoginStatus, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	// để tạm để test auth, chuyển về feature User sau
	GetUserInfo(ctx context.Context, id int64) (*getUserInfoResult, error)
}

type useCase struct {
	repo  Repository
	maker token.Maker
}

func NewUseCase(repo Repository, maker token.Maker) UseCase {
	return &useCase{
		repo:  repo,
		maker: maker,
	}
}

func (u *useCase) RegisterUser(ctx context.Context, user *User) error {
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

func (u *useCase) Login(ctx context.Context, rq LoginRequest) (*loginResult, LoginStatus, error) {
	// verify thông tin đăng nhập
	// tạo và gửi về access token và refresh token
	user, err := u.repo.GetUserByUsernamOrEmail(ctx, rq.Identifier, rq.Identifier)
	if err != nil {
		return nil, LoginStatusSystemError, fmt.Errorf("get user failed: %w", err)
	}
	if user == nil {
		return nil, LoginStatusUserNotFound, nil
	}

	err = utils.VerifyPassword(rq.Password, user.passwordHash)
	if err != nil {
		return nil, LoginStatusInvalidPassword, nil
	}

	accessPayload := token.NewAccessPayload(user.id, time.Now().Add(time.Minute*15))

	accessToken, err := u.maker.CreateToken(accessPayload)
	if err != nil {
		return nil, LoginStatusSystemError, fmt.Errorf("failed to generate accessToken: %w", err)
	}

	refreshTokenExpiresAt := time.Now().Add(24 * time.Hour * 7)
	refreshPayload := token.NewRefreshPayload(user.id, refreshTokenExpiresAt)
	refreshToken, err := u.maker.CreateToken(refreshPayload)
	if err != nil {
		return nil, LoginStatusSystemError, fmt.Errorf("failed to generate accessToken: %w", err)
	}
	// Lưu Refresh Token vào database (để quản lý và thu hồi khi cần)
	// hash token trong thực tế
	if err := u.repo.StoreRefreshToken(ctx, user.id, refreshToken, refreshTokenExpiresAt); err != nil {
		return nil, LoginStatusSystemError, fmt.Errorf("failed to save refresh token: %w", err)
	}
	response := &loginResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return response, LoginStatusSuccess, nil
}

func (u *useCase) GetUserInfo(ctx context.Context, id int64) (*getUserInfoResult, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w ", err)
	}
	if user == nil {
		return nil, Err.ErrUserNotFound
	}
	info := &getUserInfoResult{
		Username:    user.username,
		Email:       user.email,
		FullName:    user.fullName,
		PhoneNumber: user.phoneNumber,
	}
	return info, nil
}

func (u useCase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	// xác thực token, kiểm tra token trong database
	// xóa và tạo lại refresh token cũ trong db nếu sắp hết hạn
	// cấp access token mới
	payload, err := u.maker.VerifyToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to verify refresh token: %w", err)
	}

	storeToken, err := u.repo.GetRefreshToken(ctx, payload.GetUserID())
	if err != nil {
		return "", fmt.Errorf("failed to get refresh token: %w", err)
	}

	if time.Now().After(storeToken.Expires_at) {
		return "", fmt.Errorf("refresh token expired: %w", err)
	}
	if storeToken.Token != refreshToken {
		return "", fmt.Errorf("invalid refresh token")
	}

	accessPayload := token.NewAccessPayload(payload.GetUserID(), time.Now().Add(time.Minute*15))
	newAccessToken, err := u.maker.CreateToken(accessPayload)
	if err != nil {
		return "", fmt.Errorf("failed to generate accessToken: %w", err)
	}
	return newAccessToken, nil
}

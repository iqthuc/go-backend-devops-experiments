package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByUsernamOrEmail(ctx context.Context, username, email string) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
	StoreRefreshToken(ctx context.Context, userID int64, token string, exp time.Time) error
	GetRefreshToken(ctx context.Context, userID int64) (getRefreshTokenResult, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (username, email, password_hash, full_name, phone_number)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.QueryContext(
		ctx,
		query,
		user.username, //$1
		user.email,
		user.passwordHash,
		user.fullName,
		user.phoneNumber,
	)
	return err
}

func (r *repository) GetUserByUsernamOrEmail(ctx context.Context, username, email string) (*User, error) {
	query := `
		SELECT id, username, email, password_hash, full_name, phone_number, is_active, created_at, updated_at, deleted_at
		FROM users
		WHERE (username = $1 OR email = $2)
		AND deleted_at IS NULL
	`
	user := &User{}
	err := r.db.QueryRowContext(ctx, query, username, email).Scan(
		&user.id,
		&user.username,
		&user.email,
		&user.passwordHash,
		&user.fullName,
		&user.phoneNumber,
		&user.isActive,
		&user.createdAt,
		&user.updatedAt,
		&user.deletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	return user, nil
}

func (r *repository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT id, username, email, password_hash, full_name, phone_number, is_active, created_at, updated_at, deleted_at
		FROM users
		WHERE (id = $1)
		AND deleted_at IS NULL
	`
	user := &User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.id,
		&user.username,
		&user.email,
		&user.passwordHash,
		&user.fullName,
		&user.phoneNumber,
		&user.isActive,
		&user.createdAt,
		&user.updatedAt,
		&user.deletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	return user, nil
}

func (r repository) StoreRefreshToken(ctx context.Context, userID int64, token string, exp time.Time) error {
	query := `
 		INSERT INTO refresh_tokens (user_id, token, expires_at)
        VALUES ($1, $2, $3)
        ON CONFLICT (user_id)
        DO UPDATE SET token = EXCLUDED.token, expires_at = EXCLUDED.expires_at
	`
	_, err := r.db.ExecContext(ctx, query, userID, token, time.Now().Add(24*time.Hour*7))
	return err
}

func (r repository) GetRefreshToken(ctx context.Context, userID int64) (getRefreshTokenResult, error) {
	query := `SELECT user_id, token, expires_at FROM refresh_token WHERE user_id = $1`
	token := getRefreshTokenResult{}
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&token.UserID, &token.Token, &token.Expires_at)
	if err != nil {
		return getRefreshTokenResult{}, fmt.Errorf("database query error: %w", err)
	}
	return token, nil
}

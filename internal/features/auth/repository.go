package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByUsernamOrEmail(ctx context.Context, username, email string) (*User, error)
	GetUserByID(ctx context.Context, id int64) (*User, error)
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

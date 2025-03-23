package auth

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
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

func (r *repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	query := `
		SELECT id, username, email, password_hash, full_name, phone_number, is_active, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1
	`
	var user User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
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
	if err != nil {
		return User{}, fmt.Errorf("database query error: %w", err)
	}
	return User{}, nil
}

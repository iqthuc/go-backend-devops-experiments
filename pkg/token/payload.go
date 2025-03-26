package token

import (
	"time"

	"github.com/iqthuc/sport-shop/utils"
)

type TokenType int

const (
	Access TokenType = iota
	Refresh
)

type Payload interface {
	IsValid() error
	GetUserID() int64
	GetTokenType() TokenType
}

type RegistersClaims struct {
	ExpiresAt time.Time `json:"expires_at"`
	IssuedAt  time.Time `json:"issued_at"`
	UserId    int64     `json:"user_id"`
}

type BasePayload struct {
	RegistersClaims
	TokenType TokenType `json:"token_type"`
}

// // Access payload
type AccessPayload struct {
	BasePayload
}

func (p *AccessPayload) IsValid() error {
	if time.Now().After(p.ExpiresAt) {
		return Errors.ExpiredToken
	}
	return nil
}
func (p *AccessPayload) GetUserID() int64 {
	return p.UserId
}

func (p *AccessPayload) GetTokenType() TokenType {
	return p.TokenType
}

func NewAccessPayload(userID int64, expiresAt time.Time) Payload {
	return &AccessPayload{
		BasePayload: BasePayload{
			RegistersClaims: RegistersClaims{
				ExpiresAt: expiresAt,
				IssuedAt:  time.Now(),
				UserId:    userID,
			},
			TokenType: Access,
		},
	}
}

// // Refresh payload
type RefreshPayload struct {
	BasePayload
	SessionID string `json:"session_id,omitempty"`
	DeviceID  string `json:"device_id"`
}

func (p *RefreshPayload) IsValid() error {
	if time.Now().After(p.ExpiresAt) {
		return Errors.ExpiredToken
	}
	return nil
}
func (p *RefreshPayload) GetUserID() int64 {
	return p.UserId
}
func (p *RefreshPayload) GetTokenType() TokenType {
	return p.TokenType
}
func NewRefreshPayload(userID int64, expiresAt time.Time) Payload {
	return &RefreshPayload{
		BasePayload: BasePayload{
			RegistersClaims: RegistersClaims{
				ExpiresAt: expiresAt,
				IssuedAt:  time.Now(),
				UserId:    userID,
			},
			TokenType: Refresh,
		},

		SessionID: utils.GenerateSecretKey(),
	}
}

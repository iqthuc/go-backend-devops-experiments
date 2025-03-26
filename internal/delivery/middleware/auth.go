package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/iqthuc/sport-shop/pkg/token"
	"github.com/iqthuc/sport-shop/utils"
)

type AuthMiddleware struct {
	maker token.Maker
}

func NewAuthMiddleware(maker token.Maker) *AuthMiddleware {
	return &AuthMiddleware{maker: maker}
}

func (a *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorJsonResponse(w, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.ErrorJsonResponse(w, http.StatusUnauthorized, "Invalid authorization format")
			return
		}

		claims, err := a.maker.VerifyToken(tokenParts[1])
		if claims.GetTokenType() != token.Access {
			utils.ErrorJsonResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		if err != nil {
			if errors.Is(err, token.Errors.ExpiredToken) {
				utils.ErrorJsonResponse(w, http.StatusUnauthorized, "Expired token")
				return
			}
			utils.ErrorJsonResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		ctx := context.WithValue(r.Context(), ContextKey.UserID, claims.GetUserID())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

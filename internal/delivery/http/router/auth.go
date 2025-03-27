package router

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/iqthuc/sport-shop/config"
	"github.com/iqthuc/sport-shop/internal/delivery/http/middleware"
	"github.com/iqthuc/sport-shop/internal/features/auth"
	"github.com/iqthuc/sport-shop/pkg/token"
)

func IntAuthRouter(r *http.ServeMux, db *sql.DB) {
	secretKey := os.Getenv(config.Env.SecretKey)
	maker := token.NewJWTMaker(secretKey)

	authRepo := auth.NewRepository(db)
	authUseCase := auth.NewUserCase(authRepo, maker)
	authHandler := auth.NewHandler(authUseCase)

	authMW := middleware.NewAuthenticator(maker)

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /register", authHandler.RegisterUser)
	authRouter.HandleFunc("POST /login", authHandler.Login)
	authRouter.HandleFunc("POST /refresh_token", authHandler.RefreshToken)
	authRouter.HandleFunc("GET /profile", middleware.Wrap(authHandler.GetUserInfo, authMW.Auth))

	r.Handle("/auth/", http.StripPrefix("/auth", authRouter))
}

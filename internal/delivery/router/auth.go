package router

import (
	"database/sql"
	"net/http"

	"github.com/iqthuc/sport-shop/internal/delivery/middleware"
	"github.com/iqthuc/sport-shop/internal/features/auth"
	"github.com/iqthuc/sport-shop/pkg/token"
)

func IntAuthRouter(r *http.ServeMux, db *sql.DB, maker token.Maker) {
	authMiddleware := middleware.NewAuthMiddleware(maker)

	authRepo := auth.NewRepository(db)
	authUseCase := auth.NewUserCase(authRepo, maker)
	authHandler := auth.NewHandler(authUseCase)

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /register", authHandler.RegisterUser)
	authRouter.HandleFunc("POST /login", authHandler.Login)
	authRouter.HandleFunc("GET /profile", middleware.Wrap(authHandler.GetUserInfo, authMiddleware.Auth))
	middlewares := middleware.Apply(middleware.Logger)
	r.Handle("/auth/", http.StripPrefix("/auth", middlewares(authRouter)))
}

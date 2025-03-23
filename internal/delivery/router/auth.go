package router

import (
	"database/sql"
	"net/http"

	"github.com/iqthuc/sport-shop/internal/delivery/middleware"
	"github.com/iqthuc/sport-shop/internal/features/auth"
)

func IntAuthRouter(r *http.ServeMux, db *sql.DB) {
	authRepo := auth.NewRepository(db)
	authUseCase := auth.NewUserCase(authRepo)
	authHandler := auth.NewHandler(authUseCase)

	authRouter := http.NewServeMux()
	authRouter.HandleFunc("POST /register", authHandler.RegisterUser)

	middlewares := middleware.Group(middleware.Logger)
	r.Handle("/auth/", http.StripPrefix("/auth", middlewares(authRouter)))
}

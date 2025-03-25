package router

import (
	"database/sql"
	"net/http"

	"github.com/iqthuc/sport-shop/internal/delivery/middleware"
	"github.com/iqthuc/sport-shop/internal/features/admin"
)

func InitAdminRouter(r *http.ServeMux, db *sql.DB) {
	adminRepo := admin.NewRepository(db)
	adminUseCase := admin.NewUserCase(adminRepo)
	adminHandler := admin.NewHandler(adminUseCase)

	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("GET /products/{id}", adminHandler.GetProductDetails)
	adminRouter.HandleFunc("GET /products", adminHandler.GetProducts)

	middlewares := middleware.Apply(middleware.Logger)
	r.Handle("/admin/", http.StripPrefix("/admin", middlewares(adminRouter)))
}

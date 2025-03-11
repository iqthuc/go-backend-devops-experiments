package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/iqthuc/sport-shop/internal/delivery/middleware"
	"github.com/iqthuc/sport-shop/internal/features/admin"
)

func InitAdminRouter(r *http.ServeMux, db *sql.DB) {
	fmt.Println("init admin router")
	adminRepo := admin.NewRepository(db)
	adminUseCase := admin.NewService(adminRepo)
	adminHandler := admin.NewHandler(adminUseCase)

	adminRouter := http.NewServeMux()
	adminRouter.HandleFunc("GET /products", adminHandler.GetProducts)

	middlewares := middleware.Group(middleware.Logger)
	r.Handle("/admin/", http.StripPrefix("/admin", middlewares(adminRouter)))
}

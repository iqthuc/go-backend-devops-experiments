package router

import (
	"database/sql"
	"net/http"

	"github.com/iqthuc/sport-shop/internal/features/product"
)

func InitProductRouter(r *http.ServeMux, db *sql.DB) {
	productRepo := product.NewRepository(db)
	productUseCase := product.NewService(productRepo)
	productHandler := product.NewHandler(productUseCase)
	productRouter := http.NewServeMux()
	productRouter.HandleFunc("GET /", productHandler.GetProducts)
	productRouter.HandleFunc("GET /{id}", productHandler.GetProductDetails)

	r.Handle("/api/products/", http.StripPrefix("/api/products", productRouter))

}

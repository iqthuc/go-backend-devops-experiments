package router

import (
	"database/sql"
	"net/http"

	"github.com/iqthuc/sport-shop/internal/features/product"
	"github.com/redis/go-redis/v9"
)

func InitProductRouter(r *http.ServeMux, db *sql.DB, redis *redis.Client) {
	productRepo := product.NewRepository(db)
	productUseCase := product.NewUseCase(productRepo, redis)
	productHandler := product.NewHandler(productUseCase)
	productRouter := http.NewServeMux()
	productRouter.HandleFunc("GET /", productHandler.GetProducts)
	productRouter.HandleFunc("GET /{id}", productHandler.GetProductDetails)

	r.Handle("/api/products/", http.StripPrefix("/api/products", productRouter))

}

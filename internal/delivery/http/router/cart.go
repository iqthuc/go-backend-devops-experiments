package router

import (
	"net/http"

	"github.com/iqthuc/sport-shop/internal/features/cart"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitCartRouter(r *http.ServeMux, client *mongo.Client) {
	cartRepo := cart.NewRepository(client)
	cartUseCase := cart.NewUseCase(cartRepo)
	cartHandler := cart.NewHandler(cartUseCase)
	cartRouter := http.NewServeMux()
	cartRouter.HandleFunc("GET /get_cart", cartHandler.GetCart)
	cartRouter.HandleFunc("POST /add_to_cart", cartHandler.AddToCart)

	r.Handle("/api/carts/", http.StripPrefix("/api/carts", cartRouter))

}

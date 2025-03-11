package initializer

import (
	"log"
	"net/http"

	"github.com/iqthuc/sport-shop/config"
	"github.com/iqthuc/sport-shop/internal/delivery/router"
	"github.com/iqthuc/sport-shop/pkg/database"
)

func InitServer() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := http.NewServeMux()

	router.InitAdminRouter(r, db)
	router.InitProductRouter(r, db)

	http.ListenAndServe(config.SERVER_ADDRESS, r)
}

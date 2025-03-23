package initializer

import (
	"log"
	"net"
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
	router.IntAuthRouter(r, db)

	server := &http.Server{
		Addr:    config.SERVER_ADDRESS,
		Handler: r,
	}
	ln, err := net.Listen("tcp", config.SERVER_ADDRESS)
	if err != nil {
		log.Fatalf("Failed to bind address: %v", err)
	}
	log.Println("Server started successfully on", config.SERVER_ADDRESS)
	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

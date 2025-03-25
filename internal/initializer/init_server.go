package initializer

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/iqthuc/sport-shop/config"
	"github.com/iqthuc/sport-shop/internal/delivery/router"
	"github.com/iqthuc/sport-shop/pkg/database"
	"github.com/iqthuc/sport-shop/pkg/token"
)

func InitServer() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	s := http.NewServeMux()

	secretKey := os.Getenv(config.Env.SecretKey)
	maker := token.NewJWTMaker(secretKey)

	router.InitAdminRouter(s, db)
	router.InitProductRouter(s, db)
	router.IntAuthRouter(s, db, maker)

	serverAddress := os.Getenv(config.Env.ServerAddress)
	server := &http.Server{
		Addr:    serverAddress,
		Handler: s,
	}

	ln, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to bind address: %v", err)
	}
	log.Println("Server started successfully on", serverAddress)

	if err := server.Serve(ln); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

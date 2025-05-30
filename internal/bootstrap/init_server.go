package initializer

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/iqthuc/sport-shop/config"
	"github.com/iqthuc/sport-shop/internal/database"
	"github.com/iqthuc/sport-shop/internal/delivery/http/middleware"
	"github.com/iqthuc/sport-shop/internal/delivery/http/router"
	"github.com/iqthuc/sport-shop/pkg/cache"
)

const (
	RequestLimit    = 100              // Giới hạn số lượng request tối đa
	RateLimitWindow = 15 * time.Minute // Thời gian giới hạn (15 phút)
)

func InitServer() {
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	rdb, err := cache.NewRedisClient()
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := database.NewMongoDB()
	if err != nil {
		log.Fatal(err)
	}

	s := http.NewServeMux()
	router.InitAdminRouter(s, db)
	router.InitProductRouter(s, db, rdb)
	router.IntAuthRouter(s, db)
	router.IntGraphqlRouter(s, db)
	router.InitCartRouter(s, mongoClient)

	loggerMW := middleware.NewLogger()
	limiterMW := middleware.NewRateLimiter(rdb, RequestLimit, RateLimitWindow)
	middlewares := middleware.Apply(loggerMW.Logger, limiterMW.RateLimiting)

	serverAddress := os.Getenv(config.Env.ServerAddress)
	server := &http.Server{
		Addr:    serverAddress,
		Handler: middlewares(s),
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

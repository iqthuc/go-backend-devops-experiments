package utils

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/iqthuc/sport-shop/config"
)

func InitTestConfig() {
	_, basePath, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Error finding the caller's path")
	}

	envPath := filepath.Join(filepath.Dir(basePath), "../dev.env")
	err := config.LoadEnv(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

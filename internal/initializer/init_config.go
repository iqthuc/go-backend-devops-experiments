package initializer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/iqthuc/sport-shop/config"
)

func getEnvPath(filePath string) string {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return filepath.Join(rootPath, filePath)
}

func InitConfig() {
	envPath := getEnvPath("dev.env")
	err := config.LoadEnv(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

package initializer

import (
	"log"
	"os"
	"path/filepath"

	"github.com/iqthuc/sport-shop/config"
)

func InitConfig() {
	envPath := getEnvPath("dev.env")
	err := config.LoadEnv(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func getEnvPath(filePath string) string {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return filepath.Join(rootPath, filePath)
}

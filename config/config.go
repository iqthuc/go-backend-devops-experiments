package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// type Config struct {
// 	Server   ServerConfig
// 	Database DatabaseConfig
// 	Logger   LoggerConfig
// }
// type ServerConfig struct {
// 	Host string
// 	Port string
// }

// type DatabaseConfig struct {
// 	Host     string
// 	Port     string
// 	User     string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

// type LoggerConfig struct {
// 	Level string
// }

// func LoadConfig() *Config {
// 	envPath := getEnvPath("dev.env")
// 	err := LoadEnv(envPath)
// 	if err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

//		return &Config{
//			Server:   ServerConfig{},
//			Database: DatabaseConfig{},
//			Logger:   LoggerConfig{},
//		}
//	}
func getEnvPath(filePath string) string {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	return filepath.Join(rootPath, filePath)
}

func LoadEnv(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)
		os.Setenv(key, value)
	}
	return scanner.Err()
}

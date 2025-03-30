package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration settings
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// LoadConfig loads configuration from environment variables
func LoadConfig(envFile string) *Config {
	if envFile != "" {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Printf("Error loading %s file: %v", envFile, err)
		}
	}

	portStr := os.Getenv("DATABASE_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Error converting port to integer: %v\n", err)
		// Set a default port
		port = 5432
	}

	config := &Config{
		DBHost:     getEnv("DATABASE_HOST", "localhost"),
		DBPort:     port,
		DBUser:     getEnv("DATABASE_USERNAME", "postgres"),
		DBPassword: getEnv("DATABASE_PASSWORD", "postgres"),
		DBName:     getEnv("DATABASE_NAME", "taskdb"),
		ServerPort: getEnv("SERVER_PORT", "8888"),
	}

	return config
}

// Helper function to get an environment variable or a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

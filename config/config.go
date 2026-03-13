package config

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecretKey   string
}

var config *Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")

	if version == "" {
		log.Fatal("VERSION environment variable is not set")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("SERVICE_NAME environment variable is not set")
		os.Exit(1)

	}

	httpPort, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatal("HTTP_PORT environment variable is not set or is not a valid integer")
		os.Exit(1)
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable is not set")
		os.Exit(1)
	}

	config = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		JwtSecretKey: secretKey,
	}

}
func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}

package config

import (
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

type DBConfig struct{
	Host string
	Port int
	Name string
	User string
	Password string
	EnableSSLMODE bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecretKey   string
	DB             *DBConfig
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

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		log.Fatal("host is required")
		os.Exit(1)
	}

	dbport := os.Getenv("DB_PORT")
	if dbport == "" {
		log.Fatal("port is required")
		os.Exit(1)
	}
	dbprt, err := strconv.ParseInt(dbport,10,64)
	if err != nil {
		log.Fatal("db port should be an integer")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB name is required")
		os.Exit(1)
	}
	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		log.Fatal("DB User name is required")
		os.Exit(1)
	}
	dbpass := os.Getenv("DB_PASSWORD")
	if dbpass == "" {
		log.Fatal("Password is required")
		os.Exit(1)
	}
	enablessl := os.Getenv("DB_ENABLE_SSL_MODE")
	
	sslMode, err := strconv.ParseBool(enablessl)
	if err != nil{
		log.Fatal("Invalid ssl mode")
		os.Exit(1)
	}

	 dbConf := &DBConfig{
		Host: dbhost,
		Port: int(dbprt),
		Name: dbName,
		User: dbuser,
		Password: dbpass,
		EnableSSLMODE: sslMode,
	}
	

	config = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		JwtSecretKey: secretKey,
	    DB: dbConf,
	}

}
func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}

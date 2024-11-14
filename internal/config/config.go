package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	ServerAddress string
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        string
}

func LoadConfig() *Config {
	fmt.Println("Load config...")

	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBName:        getEnv("DB_NAME", "mydb"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "3306"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Using default %s for %s", defaultValue, key)
		return defaultValue
	}
	return value
}

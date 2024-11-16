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

func (cfg *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
}

func LoadConfig() *Config {
	log.Println("Load config...")

	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DBUser:        getEnv("DB_USER", "cluster_user"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBName:        getEnv("DB_NAME", "cluster_mate"),
		DBHost:        getEnv("DB_HOST", "localhost"),
		DBPort:        getEnv("DB_PORT", "3306"),
	}
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Using \"%s\" for %s", defaultValue, key)
		return defaultValue
	}
	log.Printf("Using \"%s\" for %s", defaultValue, key)
	return value
}

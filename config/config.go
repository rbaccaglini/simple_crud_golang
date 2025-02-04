package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config armazena as configurações da aplicação
type Config struct {
	Port             string
	MongoURI         string
	DatabaseName     string
	UserDbCollection string
	JWTSecret        string
}

func LoadConfig() (*Config, error) {
	godotenv.Load("../../config/.env")

	config := &Config{
		Port:             mustGetEnv("APP_PORT"),
		MongoURI:         mustGetEnv("MONGODB_URL"),
		DatabaseName:     mustGetEnv("MONGODB_DB"),
		UserDbCollection: mustGetEnv("MONGODB_USER_DB_COLLECTION"),
		JWTSecret:        mustGetEnv("JWT_SECRET_KEY"),
	}

	return config, nil
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Erro: A variável de ambiente %s não está definida", key)
	}
	return value
}

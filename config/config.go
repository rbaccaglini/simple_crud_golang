package config

import (
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

func LoadConfig() *Config {
	godotenv.Load("../../config/.env")

	config := &Config{
		Port:             getEnv("PORT", "8080"),
		MongoURI:         getEnv("MONGODB_URL", "mongodb://localhost:27017"),
		DatabaseName:     getEnv("MONGODB_DB", "crud-init"),
		UserDbCollection: getEnv("MONGODB_USER_DB_COLLECTION", "users"),
		JWTSecret:        getEnv("JWT_SECRET_KEY", ""),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

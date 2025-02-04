package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"

	user_handler "github.com/rbaccaglini/simple_crud_golang/internal/handlers/user"
	user_repository "github.com/rbaccaglini/simple_crud_golang/internal/repositories/user"
	"github.com/rbaccaglini/simple_crud_golang/internal/routes"
	user_service "github.com/rbaccaglini/simple_crud_golang/internal/services/user"
	"github.com/rbaccaglini/simple_crud_golang/pkg/database/mongodb"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
)

func main() {
	logger.Info("Starting server...")

	godotenv.Load("../../config/.env")
	logger.Info("Config env loaded")

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("Error trying to connect on DB", err)
		return
	}
	logger.Info("DB Connecteded")

	userHandler := initDependencies(database)
	router := gin.Default()
	routes.InitRouter(&router.RouterGroup, userHandler)
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func initDependencies(database *mongo.Database) user_handler.UserHandlerInterface {
	repo := user_repository.NewUserRepository(database)
	service := user_service.NewUserDomainService(repo)
	return user_handler.NewUserHandlerInterface(service)
}

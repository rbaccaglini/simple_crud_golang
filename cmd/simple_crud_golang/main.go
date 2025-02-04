package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rbaccaglini/simple_crud_golang/config"
	user_handler "github.com/rbaccaglini/simple_crud_golang/internal/handlers/user"
	user_repository "github.com/rbaccaglini/simple_crud_golang/internal/repositories/user"
	"github.com/rbaccaglini/simple_crud_golang/internal/routes"
	user_service "github.com/rbaccaglini/simple_crud_golang/internal/services/user"
	"github.com/rbaccaglini/simple_crud_golang/pkg/database/mongodb"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
)

func main() {
	logger.Info("Starting server...")

	config := config.LoadConfig()

	database, err := mongodb.NewMongoDBConnection(context.Background(), config)
	if err != nil {
		logger.Error("Error trying to connect on DB: %v", err)
		return
	}
	logger.Info("DB Connecteded")

	userHandler := initDependencies(database, config)
	router := gin.Default()
	routes.InitRouter(&router.RouterGroup, userHandler)
	if err := router.Run(fmt.Sprintf(":%s", config.Port)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func initDependencies(
	database *mongo.Database,
	config *config.Config,
) user_handler.UserHandlerInterface {
	repo := user_repository.NewUserRepository(database, config)
	service := user_service.NewUserDomainService(repo)
	return user_handler.NewUserHandlerInterface(service)
}

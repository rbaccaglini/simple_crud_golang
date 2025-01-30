package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/database/mongodb"
	"github.com/rbaccaglini/simple_crud_golang/src/controller/routes"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
)

func main() {

	logger.Info("Starting aplication")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect on DB: %v", err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	routes.InitRouter(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

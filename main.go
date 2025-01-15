package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rbaccaglini/simple_crud_golang/src/controller"
	"github.com/rbaccaglini/simple_crud_golang/src/controller/routes"
	"github.com/rbaccaglini/simple_crud_golang/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRouter(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

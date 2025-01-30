package main

import (
	"github.com/rbaccaglini/simple_crud_golang/src/controller"
	"github.com/rbaccaglini/simple_crud_golang/src/model/repository"
	"github.com/rbaccaglini/simple_crud_golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}

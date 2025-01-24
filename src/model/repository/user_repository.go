package repository

import (
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB_COLLECTION = "users"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPass(string, string) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteUser(string) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
}

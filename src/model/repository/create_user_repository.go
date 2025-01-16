package repository

import (
	"context"
	"os"

	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_DB            = "MONGODB_USER_DB"
	MONGODB_USER_DB_COLLECTION = "users"
)

var JOURNEY = zap.String("journey", "CreateUserRepository")

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", JOURNEY)

	collection_name := MONGODB_USER_DB_COLLECTION
	collection := ur.databaseConnection.Collection(os.Getenv(collection_name))

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetId(result.InsertedID.(string))

	return userDomain, nil

}

package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	entity "github.com/rbaccaglini/simple_crud_golang/src/model/repository/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(id string, domain model.UserDomainInterface) *rest_err.RestErr {
	journey := zap.String("journey", "UpdateUserRepository")
	logger.Info("Init update user repository", journey)

	collection := ur.databaseConnection.Collection(os.Getenv(MONGODB_USER_DB_COLLECTION))

	value := entity.ConverterDomainToEntity(domain)
	userId, _ := primitive.ObjectIDFromHex(id)

	var values = bson.D{}

	if value.Name != "" {
		values = append(values, bson.E{Key: "name", Value: value.Name})
	}

	if value.Age != 0 {
		values = append(values, bson.E{Key: "age", Value: value.Age})
	}

	update := bson.D{
		{
			Key: "$set", Value: values,
		},
	}

	_, err := collection.UpdateByID(context.Background(), userId, update)
	if err != nil {
		errMessage := fmt.Sprintf("Error on update user with id %s", id)
		logger.Error(errMessage, err, journey)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(fmt.Sprintf("User (id: %s) updated with success", id), journey)

	return nil
}

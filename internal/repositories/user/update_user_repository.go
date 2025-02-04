package user_repository

import (
	"context"
	"fmt"
	"os"

	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/internal/util/converter"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(user domain.UserDomainInterface, uid string) *rest_err.RestErr {
	journey := zap.String("journey", "UpdateUserRepository")
	logger.Info("Init update user repository", journey)

	collection := ur.databaseConnection.Collection(os.Getenv(MONGODB_USER_DB_COLLECTION))

	value := converter.ConverterDomainToEntity(user)
	userId, errId := primitive.ObjectIDFromHex(uid)
	if errId != nil {
		logger.Error("invalid user id", errId, journey)
		return rest_err.NewInternalServerError(fmt.Sprintf("invalid user id: %s", errId.Error()))
	}

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
		errMessage := fmt.Sprintf("error on update user with id %s", uid)
		logger.Error(errMessage, err, journey)
		return rest_err.NewInternalServerError(
			fmt.Sprintf("error on update user with id %s: %s", uid, err.Error()))
	}

	logger.Info(fmt.Sprintf("User (id: %s) updated with success", uid), journey)

	return nil
}

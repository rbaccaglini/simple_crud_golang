package user_repository

import (
	"context"
	"fmt"

	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/entity"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) InsertUser(user domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateUserRepository")
	logger.Info("Init createUser repository", journey)

	collection := ur.databaseConnection.Collection(ur.config.UserDbCollection)

	value := entity.ConverterDomainToEntity(user)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return entity.ConverterEntityToDomain(*value), nil
}

func (ur *userRepository) DeleteUser(uid string) *rest_err.RestErr {
	journey := zap.String("journey", "DeleteUserRepository")
	logger.Info("Init createUser repository", journey)

	puid, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		logger.Error("create user repository", err, journey)
		return rest_err.NewInternalServerError(err.Error())
	}
	filter := bson.D{{Key: "_id", Value: puid}}

	collection := ur.databaseConnection.Collection(ur.config.UserDbCollection)
	if _, errDel := collection.DeleteOne(context.Background(), filter); errDel != nil {
		logger.Error("create user repository", errDel, journey)
		return rest_err.NewInternalServerError(errDel.Error())
	}
	return nil
}

func (ur *userRepository) UpdateUser(user domain.UserDomainInterface, uid string) *rest_err.RestErr {
	journey := zap.String("journey", "UpdateUserRepository")
	logger.Info("Init update user repository", journey)

	collection := ur.databaseConnection.Collection(ur.config.UserDbCollection)

	value := entity.ConverterDomainToEntity(user)
	userId, _ := primitive.ObjectIDFromHex(uid)

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
		errMessage := fmt.Sprintf("Error on update user with id %s", uid)
		logger.Error(errMessage, err, journey)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(fmt.Sprintf("User (id: %s) updated with success", uid), journey)

	return nil
}

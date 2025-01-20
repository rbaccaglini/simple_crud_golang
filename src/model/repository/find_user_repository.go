package repository

import (
	"context"
	"fmt"

	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	entity "github.com/rbaccaglini/simple_crud_golang/src/model/repository/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "FindUserByEmail"))
	userFound, err := ur.findBy("email", email)
	return userFound, err
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "FindUserByEmail"))
	userFound, err := ur.findBy("id", id)
	return userFound, err
}

func (ur *userRepository) FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindAllUsers")
	logger.Info("Init Find All Users repository", journey)

	filter := bson.D{}
	collection := ur.databaseConnection.Collection(MONGODB_USER_DB_COLLECTION)
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		errMessage := "Error trying to get all users"
		logger.Error(errMessage, err, journey)
		return nil, rest_err.NewInternalServerError(errMessage)
	}

	resp := []model.UserDomainInterface{}
	for cur.Next(context.Background()) {
		userEntity := &entity.UserEntity{}
		err := cur.Decode(userEntity)
		if err != nil {
			errMessage := "Error trying to get all users"
			logger.Error(errMessage, err, journey)
			return nil, rest_err.NewInternalServerError(errMessage)
		}
		userDomain := entity.ConverterEntityToDomain(*userEntity)
		resp = append(resp, userDomain)
	}

	return resp, nil
}

func (ur *userRepository) findBy(field string, value string) (model.UserDomainInterface, *rest_err.RestErr) {

	journey := zap.String("journey", "FindUserByField")

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: field, Value: value}}
	if field == "id" {
		field = "_id"
		valueId, _ := primitive.ObjectIDFromHex(value)
		filter = bson.D{{Key: field, Value: valueId}}
	}

	collection := ur.databaseConnection.Collection(MONGODB_USER_DB_COLLECTION)

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errMessage := fmt.Sprintf("User not found with %s %s", field, value)
			logger.Error(errMessage, err, journey)
			return nil, rest_err.NewNotFoundError(errMessage)
		}

		errMessage := fmt.Sprintf("Error trying to find user by %s", field)
		logger.Error(errMessage, err, journey)

		return nil, rest_err.NewInternalServerError(errMessage)
	}

	return entity.ConverterEntityToDomain(*userEntity), nil
}

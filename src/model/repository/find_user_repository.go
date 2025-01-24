package repository

import (
	"context"
	"encoding/json"
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
	userFound, err := ur.findBy(bson.D{{Key: "email", Value: email}})
	return userFound, err
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "FindUserByEmail"))
	userFound, err := ur.findBy(bson.D{{Key: "id", Value: id}})
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

func (ur *userRepository) FindUserByEmailAndPass(email string, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository", zap.String("journey", "FindUserByEmailAndPass"))

	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	userFound, err := ur.findBy(filter)
	return userFound, err
}

func (ur *userRepository) findBy(filter bson.D) (model.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindUserByField")

	userEntity := &entity.UserEntity{}
	filterStg := filterAjustment(&filter)

	collection := ur.databaseConnection.Collection(MONGODB_USER_DB_COLLECTION)

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errMessage := fmt.Sprintf("User not found. Filter: %s", filterStg)
			logger.Error(errMessage, err, journey)
			return nil, rest_err.NewNotFoundError("User not found.")
		}

		errMessage := fmt.Sprintf("Error trying to find user. Filter: %s", filterStg)
		logger.Error(errMessage, err, journey)

		return nil, rest_err.NewInternalServerError("Error trying to find user")
	}

	return entity.ConverterEntityToDomain(*userEntity), nil
}

func converterFilterByJson(filter bson.D) string {
	journey := zap.String("journey", "FindUserByField")
	filterStg := ""
	bsonBytes, err := bson.Marshal(filter)
	if err != nil {
		logger.Error("Error on convert bson.D: %v", err, journey)
	} else {
		var bsonMap bson.M
		err := bson.Unmarshal(bsonBytes, &bsonMap)
		if err != nil {
			logger.Error("Error on convert bson.M: %v", err)
		} else {
			jsonData, err := json.Marshal(bsonMap)
			if err != nil {
				logger.Error("Error on convert bson.M to JSON: %v", err)
			} else {
				filterStg = string(jsonData)
			}
		}
	}
	return filterStg
}

func filterAjustment(filter *bson.D) string {
	filterToPrint := bson.D{}

	f := *filter

	for ix, v := range f {
		switch f[ix].Key {
		case "id":
			filterToPrint = append(filterToPrint, bson.E{Key: f[ix].Key, Value: f[ix].Value})
			f[ix].Key = "_id"
			valueId, _ := primitive.ObjectIDFromHex(v.Value.(string))
			f[ix].Value = valueId
		case "password":
			filterToPrint = append(filterToPrint, bson.E{Key: f[ix].Key, Value: "######"})
		default:
			filterToPrint = append(filterToPrint, bson.E{Key: f[ix].Key, Value: f[ix].Value})
		}
	}

	return converterFilterByJson(filterToPrint)
}

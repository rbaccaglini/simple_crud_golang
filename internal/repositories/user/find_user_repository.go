package user_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/entity"
	"github.com/rbaccaglini/simple_crud_golang/internal/util/converter"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) GetUsers() ([]domain.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindAllUsers")
	logger.Info("Init Find All Users repository", journey)

	filter := bson.D{}
	collection := ur.databaseConnection.Collection(os.Getenv(MONGODB_USER_DB_COLLECTION))
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		errMessage := "Error trying to get all users"
		logger.Error(errMessage, err, journey)
		return nil, rest_err.NewInternalServerError(errMessage)
	}
	defer cur.Close(context.Background())

	resp := []domain.UserDomainInterface{}
	for cur.Next(context.Background()) {
		userEntity := &entity.UserEntity{}
		err := cur.Decode(userEntity)
		if err != nil {
			errMessage := "Error trying to get all users"
			logger.Error(errMessage, err, journey)
			return nil, rest_err.NewInternalServerError(errMessage)
		}
		userDomain := converter.ConverterEntityToDomain(*userEntity)
		resp = append(resp, userDomain)
	}

	return resp, nil
}

func (ur *userRepository) GetUserById(uid string) (domain.UserDomainInterface, *rest_err.RestErr) {
	return ur.findBy(bson.D{{Key: "id", Value: uid}})
}

func (ur *userRepository) GetUserByEmail(email string) (domain.UserDomainInterface, *rest_err.RestErr) {
	return ur.findBy(bson.D{{Key: "email", Value: email}})
}

func (ur *userRepository) ValidateCredentials(email, password string) (domain.UserDomainInterface, *rest_err.RestErr) {

	ud := domain.NewLoginDomain(email, password)
	ud.EncryptPassword()

	filter := bson.D{{Key: "email", Value: email}, {Key: "password", Value: ud.GetPassword()}}

	u, err := ur.findBy(filter)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (ur *userRepository) findBy(filter bson.D) (domain.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "FindUserByField")

	userEntity := &entity.UserEntity{}
	filterStg := filterAjustment(&filter)

	collection := ur.databaseConnection.Collection(os.Getenv(MONGODB_USER_DB_COLLECTION))

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

	return converter.ConverterEntityToDomain(*userEntity), nil
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

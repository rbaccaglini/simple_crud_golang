package repository

import (
	"context"
	"fmt"

	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init deleteUser repository", zap.String("journey", "DeleteUserRepository"))

	collection := ur.databaseConnection.Collection(MONGODB_USER_DB_COLLECTION)

	userId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: userId}}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		errMessage := fmt.Sprintf("Error on delete user with id %s", id)
		logger.Error(errMessage, err, zap.String("journey", "DeleteUserRepository"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("User deleted with success", zap.String("journey", "DeleteUserRepository"))

	return nil
}

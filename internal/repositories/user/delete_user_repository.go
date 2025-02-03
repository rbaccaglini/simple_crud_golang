package user_repository

import (
	"context"
	"fmt"

	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(uid string) *rest_err.RestErr {
	journey := zap.String("journey", "DeleteUserRepository")
	logger.Info("Init createUser repository", journey)

	puid, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		logger.Error("invalid user id", err, journey)
		return rest_err.NewInternalServerError(fmt.Sprintf("invalid user id: %s", err.Error()))
	}
	filter := bson.D{{Key: "_id", Value: puid}}

	collection := ur.databaseConnection.Collection(ur.config.UserDbCollection)
	if _, errDel := collection.DeleteOne(context.Background(), filter); errDel != nil {
		logger.Error("error on delete user", errDel, journey)
		return rest_err.NewInternalServerError(fmt.Sprintf("error on delete user: %s", errDel.Error()))
	}
	return nil
}

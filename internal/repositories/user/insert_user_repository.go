package user_repository

import (
	"context"

	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/internal/util/converter"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) InsertUser(user domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateUserRepository")
	logger.Info("Init createUser repository", journey)

	collection := ur.databaseConnection.Collection(ur.config.UserDbCollection)

	value := converter.ConverterDomainToEntity(user)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConverterEntityToDomain(*value), nil
}

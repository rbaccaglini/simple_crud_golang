package service

import (
	"net/http"

	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Creating user", zap.String("journey", "CreateUser"))

	user, err := ud.FindUserByEmailService(userDomain.GetEmail())
	if err != nil {
		if err.Code == http.StatusNotFound {
			logger.Error("Email is already registered", err, zap.String("journey", "CreateUser"))
			return nil, rest_err.NewBadRequestError("Email is already registered")
		}
		logger.Error("Error on verify email", err, zap.String("journey", "CreateUser"))
		return nil, rest_err.NewInternalServerError("Error on verify email")
	}

	if user != nil {
		logger.Error("Email is already registered", err, zap.String("journey", "CreateUser"))
		return nil, rest_err.NewBadRequestError("Email is already registered")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepo.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}

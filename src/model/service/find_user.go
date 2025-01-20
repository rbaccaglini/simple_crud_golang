package service

import (
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user", zap.String("journey", "FindUserById"))
	return ud.userRepo.FindUserById(id)
}

func (ud *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user", zap.String("journey", "FindUserByEmail"))
	return ud.userRepo.FindUserByEmail(email)
}

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

func (ud *userDomainService) FindAllUsersService() ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding all users", zap.String("journey", "FindAllUsers"))
	return ud.userRepo.FindAllUsers()
}

func (ud *userDomainService) findUserByEmailAndPassService(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user", zap.String("journey", "FindUserByEmail"))
	return ud.userRepo.FindUserByEmailAndPass(email, password)
}

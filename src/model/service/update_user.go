package service

import (
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	journey := zap.String("journey", "UpdateUser")
	logger.Info("Creating user", journey)

	err := ud.userRepo.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error on call update user repository", err, journey)
		return err
	}

	return nil
}

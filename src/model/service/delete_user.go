package service

import (
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Deleting user", zap.String("journey", "DeleteUser"))
	return ud.userRepo.DeleteUser(id)
}

package service

import (
	"fmt"

	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.uber.org/zap"
)

var JOURNEY = zap.String("journey", "CreateUser")

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Creating user [model]", JOURNEY)
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}

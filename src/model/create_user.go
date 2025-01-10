package model

import (
	"fmt"

	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"go.uber.org/zap"
)

var JOURNEY = zap.String("journey", "CreateUser")

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Creating user [model]", JOURNEY)
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}

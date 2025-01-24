package service

import (
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(
	domain model.UserDomainInterface,
) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Finding user", zap.String("journey", "Login User Service"))
	domain.EncryptPassword()

	r, err := ud.findUserByEmailAndPassService(domain.GetEmail(), domain.GetPassword())

	if err != nil {
		if err.Code == 404 {
			return nil, "", rest_err.NewForbiddenError("Invalid credentials")
		}
		return nil, "", err
	}

	t, err := r.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return r, t, nil
}

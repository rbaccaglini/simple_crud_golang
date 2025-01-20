package service

import (
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/rbaccaglini/simple_crud_golang/src/model/repository"
)

func NewUserDomainService(userRepo repository.UserRepository) UserDomainService {
	return &userDomainService{userRepo}
}

type userDomainService struct {
	userRepo repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr

	FindUserByIdService(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(string) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteUser(string) *rest_err.RestErr
}

package user_service

import (
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	user_repository "github.com/rbaccaglini/simple_crud_golang/internal/repositories/user"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
)

func NewUserDomainService(userRepo user_repository.UserRepository) UserDomainService {
	return &userDomainService{userRepo}
}

type userDomainService struct {
	userRepo user_repository.UserRepository
}

type UserDomainService interface {
	FindAllUser() ([]domain.UserDomainInterface, *rest_err.RestErr)
	FindUserById(uid string) (domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(uid string) (domain.UserDomainInterface, *rest_err.RestErr)

	CreateUser(user domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(uid string) *rest_err.RestErr
	UpdateUser(user domain.UserDomainInterface) *rest_err.RestErr

	Login(email, password string) (string, domain.UserDomainInterface, *rest_err.RestErr)
}

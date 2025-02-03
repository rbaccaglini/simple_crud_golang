package user_service

import (
	"net/http"

	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.uber.org/zap"
)

func (us *userDomainService) FindAllUser() ([]domain.UserDomainInterface, *rest_err.RestErr) {
	return us.userRepo.GetUsers()
}

func (us *userDomainService) FindUserById(uid string) (domain.UserDomainInterface, *rest_err.RestErr) {
	return us.userRepo.GetUserById(uid)
}

func (us *userDomainService) FindUserByEmail(email string) (domain.UserDomainInterface, *rest_err.RestErr) {
	return us.userRepo.GetUserByEmail(email)
}

func (us *userDomainService) CreateUser(user domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "CreateUser")
	logger.Info("Creating user", journey)

	/* email must be unic */
	u, err := us.userRepo.GetUserByEmail(user.GetEmail())
	if err != nil {
		if err.Code != http.StatusNotFound {
			logger.Error("Email is already registered", err, zap.String("journey", "CreateUser"))
			return nil, rest_err.NewBadRequestError("Email is already registered")
		}
	}

	if u != nil {
		logger.Error("Email is already registered", err, zap.String("journey", "CreateUser"))
		return nil, rest_err.NewBadRequestError("Email is already registered")
	}

	user.EncryptPassword()

	return us.userRepo.InsertUser(user)
}

func (us *userDomainService) DeleteUser(uid string) *rest_err.RestErr {
	journey := zap.String("journey", "deleteUserService")
	logger.Info("delete user", journey)

	return us.userRepo.DeleteUser(uid)
}

func (us *userDomainService) UpdateUser(user domain.UserDomainInterface) *rest_err.RestErr {
	journey := zap.String("journey", "updateUserService")
	logger.Info("update user", journey)

	/** get current user data */
	cUser, err := us.userRepo.GetUserById(user.GetID())
	if err != nil {
		logger.Error("User not found", err, journey)
		rest_err.NewNotFoundError("User not found")
	}

	/** Check if there is a change to be made */
	if cUser.GetEmail() == user.GetEmail() && cUser.GetAge() == user.GetAge() {
		logger.Info("There is no update to do", journey)
		return nil
	}

	return us.userRepo.UpdateUser(user, user.GetID())
}

func (us *userDomainService) Login(email, password string) (string, domain.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "LoginService")
	logger.Info("login user", journey)
	u, err := us.userRepo.ValidateCredentials(email, password)
	if err != nil {
		logger.Info("Erro on validade credentians", journey)
		return "", nil, rest_err.NewForbiddenError("Invalid credentials")
	}
	t, err := u.TokenGenerate()
	if err != nil {
		logger.Info("Erro on generate token", journey)
		return "", nil, rest_err.NewInternalServerError("Error on generate token")
	}

	return t, u, nil
}

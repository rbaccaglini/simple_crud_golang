package user_service

import (
	"fmt"
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
			errMsg := "email is already registered"
			logger.Error(errMsg, err, zap.String("journey", "CreateUser"))
			return nil, rest_err.NewBadRequestError(fmt.Sprintf("%s: %s", errMsg, err.Message))
		}
	}

	if u != nil {
		errMsg := "email is already registered"
		logger.Error(errMsg, rest_err.NewBadRequestError(errMsg), zap.String("journey", "CreateUser"))
		return nil, rest_err.NewBadRequestError(errMsg)
	}

	user.EncryptPassword()

	return us.userRepo.InsertUser(user)
}

func (us *userDomainService) DeleteUser(uid string) *rest_err.RestErr {
	journey := zap.String("journey", "deleteUserService")
	logger.Info("delete user", journey)

	return us.userRepo.DeleteUser(uid)
}

func (us *userDomainService) UpdateUser(user domain.UserDomainInterface, uid string) *rest_err.RestErr {
	journey := zap.String("journey", "updateUserService")
	logger.Info("update user", journey)

	/** get current user data */
	cUser, err := us.userRepo.GetUserById(uid)
	if err != nil {
		logger.Error("user not found", err, journey)
		return rest_err.NewNotFoundError("user not found")
	}

	/** Check if there is a change to be made */

	logger.Info(fmt.Sprintf("name: '%s' | '%s'", cUser.GetName(), user.GetName()), journey)
	logger.Info(fmt.Sprintf("age: '%d' | '%d'", cUser.GetAge(), user.GetAge()), journey)

	if cUser.GetName() == user.GetName() && cUser.GetAge() == user.GetAge() {
		logger.Info("there is no update to do", journey)
		return nil
	}

	return us.userRepo.UpdateUser(user, uid)
}

func (us *userDomainService) Login(email, password string) (string, domain.UserDomainInterface, *rest_err.RestErr) {
	journey := zap.String("journey", "LoginService")
	logger.Info("login user", journey)
	u, err := us.userRepo.ValidateCredentials(email, password)
	if err != nil {
		logger.Info("error on validade credentians", journey)
		return "", nil, rest_err.NewForbiddenError("invalid credentials")
	}
	t, err := u.TokenGenerate()
	if err != nil {
		logger.Info("error on generate token", journey)
		return "", nil, rest_err.NewInternalServerError("error on generate token")
	}

	return t, u, nil
}

package user_repository

import (
	"github.com/rbaccaglini/simple_crud_golang/config"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database, config *config.Config) UserRepository {
	return &userRepository{
		database,
		config,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
	config             *config.Config
}

type UserRepository interface {
	GetUsers() ([]domain.UserDomainInterface, *rest_err.RestErr)
	GetUserById(uid string) (domain.UserDomainInterface, *rest_err.RestErr)
	GetUserByEmail(email string) (domain.UserDomainInterface, *rest_err.RestErr)

	InsertUser(user domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(uid string) *rest_err.RestErr
	UpdateUser(user domain.UserDomainInterface, uid string) *rest_err.RestErr

	ValidateCredentials(email, password string) (domain.UserDomainInterface, *rest_err.RestErr)
}

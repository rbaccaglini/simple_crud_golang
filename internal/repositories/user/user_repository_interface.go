package user_repository

import (
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB_COLLECTION = "MONGODB_USER_DB_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
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

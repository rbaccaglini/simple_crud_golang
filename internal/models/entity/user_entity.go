package entity

import (
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int8               `bson:"age,omitempty"`
}

func ConverterDomainToEntity(domain domain.UserDomainInterface) *UserEntity {
	return &UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

func ConverterEntityToDomain(entity UserEntity) domain.UserDomainInterface {
	domain := domain.NewUserDomain(entity.Email, entity.Password, entity.Name, entity.Age)
	domain.SetID(entity.ID.Hex())
	return domain
}

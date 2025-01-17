package entity

import (
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID       primitive.ObjectID `bson: "_id, omitempty"`
	Email    string             `bson: "email"`
	Password string             `bson: "password"`
	Name     string             `bson: "name"`
	Age      int8               `bson: "age"`
}

func ConverterDomainToEntity(domain model.UserDomainInterface) *UserEntity {
	return &UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

func ConverterEntityToDomain(entity UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Password, entity.Name, entity.Age)
	domain.SetId(entity.ID.Hex())
	return domain
}

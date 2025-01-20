package entity

import (
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Age      int8               `bson:"age,omitempty"`
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
	domain.SetID(entity.ID.Hex())
	return domain
}

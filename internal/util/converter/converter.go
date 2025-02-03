package converter

import (
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/entity"
	user_response "github.com/rbaccaglini/simple_crud_golang/internal/models/response/user"
)

func ConverterDomainToEntity(domain domain.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

func ConverterEntityToDomain(entity entity.UserEntity) domain.UserDomainInterface {
	domain := domain.NewUserDomain(entity.Email, entity.Password, entity.Name, entity.Age)
	domain.SetID(entity.ID.Hex())
	return domain
}

func ConvertDomainToResponse(ud domain.UserDomainInterface) user_response.UserResponse {
	return user_response.UserResponse{
		ID:    ud.GetID(),
		Email: ud.GetEmail(),
		Name:  ud.GetName(),
		Age:   ud.GetAge(),
	}
}

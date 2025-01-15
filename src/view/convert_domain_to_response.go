package view

import (
	"github.com/rbaccaglini/simple_crud_golang/src/controller/model/response"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

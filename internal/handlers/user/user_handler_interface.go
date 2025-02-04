package user_handler

import (
	"github.com/gin-gonic/gin"
	user_service "github.com/rbaccaglini/simple_crud_golang/internal/services/user"
)

func NewUserHandlerInterface(si user_service.UserDomainService) UserHandlerInterface {
	return &userHandlerInterface{
		service: si,
	}
}

type UserHandlerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindAllUser(c *gin.Context)

	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)

	Login(c *gin.Context)
}

type userHandlerInterface struct {
	service user_service.UserDomainService
}

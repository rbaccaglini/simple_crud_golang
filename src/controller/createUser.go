package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/validation"
	request "github.com/rbaccaglini/simple_crud_golang/src/controller/model/request"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
	"github.com/rbaccaglini/simple_crud_golang/src/view"
	"go.uber.org/zap"
)

var (
	JOURNEY             = zap.String("journey", "CreateUser")
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {

	logger.Info("CreateUser function called", JOURNEY)

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("There are some incorrect fields", err, JOURNEY)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		errRest := validation.ValidateUserError(err)
		c.JSON(restErr.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		logger.Error("Error creating user", err, JOURNEY)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User created successfully: %s", domain.GetName()), JOURNEY)

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domain))

}

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/validation"
	request "github.com/rbaccaglini/simple_crud_golang/src/controller/model/request"
	response "github.com/rbaccaglini/simple_crud_golang/src/controller/model/response"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"go.uber.org/zap"
)

var JOURNEY = zap.String("journey", "CreateUser")

func CreateUser(c *gin.Context) {

	logger.Info("CreateUser function called", JOURNEY)

	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("There are some incorrect fields", err, JOURNEY)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		errRest := validation.ValidateUserError(err)
		c.JSON(restErr.Code, errRest)
		return
	}

	var UserResponse response.UserResponse
	UserResponse.ID = "123"
	UserResponse.Email = UserRequest.Email
	UserResponse.Name = UserRequest.Name
	UserResponse.Age = UserRequest.Age

	logger.Info("User created successfully", JOURNEY)

	c.JSON(200, UserResponse)

}

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/validation"
	request "github.com/rbaccaglini/simple_crud_golang/src/controller/model/request"
	response "github.com/rbaccaglini/simple_crud_golang/src/controller/model/response"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		errRest := validation.ValidateUserError(err)
		c.JSON(restErr.Code, errRest)
		return
	}

	fmt.Println(UserRequest)

	var UserResponse response.UserResponse
	UserResponse.ID = "123"
	UserResponse.Email = UserRequest.Email
	UserResponse.Name = UserRequest.Name
	UserResponse.Age = UserRequest.Age
	c.JSON(200, UserResponse)

}

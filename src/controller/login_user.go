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

func (uc *userControllerInterface) Login(c *gin.Context) {
	var JOURNEY = zap.String("journey", "Login")

	logger.Info("Login function called", JOURNEY)

	var LoginRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		logger.Error("There are some incorrect fields", err, JOURNEY)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		errRest := validation.ValidateUserError(err)
		c.JSON(restErr.Code, errRest)
		return
	}

	loginDomain := model.NewLoginDomain(
		LoginRequest.Email,
		LoginRequest.Password,
	)

	domainResult, err := uc.service.LoginUserService(loginDomain)
	if err != nil {
		logger.Error("Error on call login service", err, JOURNEY)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"Logged with success",
		zap.String("userId", domainResult.GetID()),
		JOURNEY,
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}

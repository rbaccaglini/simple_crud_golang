package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init Find User controller", zap.String("journey", "FindUserById"))
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {

		logger.Error(
			"Invalid user id",
			err,
			zap.String("journey", "FindUserById"),
		)

		errorMessage := rest_err.NewBadRequestError(
			"Invalid user id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdService(userId)
	if err != nil {
		logger.Error(
			"Error on call find user by id service",
			err,
			zap.String("journey", "FindUserById"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User found by id with soccess", zap.String("journey", "FindUserById"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init Find User controller", zap.String("journey", "FindUserByEmail"))
	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error(
			"Invalid user email",
			err,
			zap.String("journey", "FindUserByEmail"),
		)

		errorMessage := rest_err.NewBadRequestError(
			"Invalid user email",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(userEmail)
	if err != nil {
		logger.Error(
			"Error on call find user by email service",
			err,
			zap.String("journey", "FindUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User found by email with soccess", zap.String("journey", "FindUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

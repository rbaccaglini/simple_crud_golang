package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/controller/model/response"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
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

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	journey := zap.String("journey", "FindAllUsers")
	logger.Info("Init Find All Users controller", journey)

	t := c.GetHeader("Authorization")
	_, err := model.VerifyToken(t)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	listUserDomain, err := uc.service.FindAllUsersService()
	if err != nil {
		logger.Error("Error on call find all users service", err, journey)
		c.JSON(err.Code, err)
		return
	}

	listView := []response.UserResponse{}
	for _, user := range listUserDomain {
		viewUser := view.ConvertDomainToResponse(user)
		listView = append(listView, viewUser)
	}

	logger.Info(fmt.Sprintf("There are %d users in DB", len(listView)), journey)

	c.JSON(http.StatusOK, listView)
}

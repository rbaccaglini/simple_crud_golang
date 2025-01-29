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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	journey := zap.String("journey", "UpdateUser")
	logger.Info("Init update user controller", journey)
	userId := c.Param("userId")

	var userRequest request.UserUpdateRequest

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error(
			"Invalid user id",
			err,
			journey,
		)
		errorMessage := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errMessage := fmt.Sprintf("There are some incorrect fields, error=%s", err.Error())
		logger.Error(errMessage, err, journey)
		restErr := rest_err.NewBadRequestError(errMessage)
		errRest := validation.ValidateUserError(err)
		c.JSON(restErr.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.UpdateUserService(userId, domain); err != nil {
		logger.Error("Error on call update user service", err, journey)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User with id %s updated successfully", userId), journey)
	c.JSON(http.StatusNoContent, nil)
}

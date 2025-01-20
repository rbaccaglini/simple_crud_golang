package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/configuration/rest_err"
	"github.com/rbaccaglini/simple_crud_golang/src/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	journey := zap.String("journey", "DeleteUser")
	logger.Info("Init delete user controller", journey)
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error(
			"Invalid user id",
			err,
			journey,
		)

		errorMessage := rest_err.NewBadRequestError(
			"Invalid user id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	if err := uc.service.DeleteUser(userId); err != nil {
		logger.Error("Error on call delete user service", err, journey)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted with success", journey)

	c.JSON(http.StatusNoContent, nil)

}

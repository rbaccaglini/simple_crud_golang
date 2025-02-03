package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) DeleteUser(c *gin.Context) {
	journey := zap.String("journey", "DeleteUser")
	logger.Info("Init delete user controller", journey)

	uid := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(uid); err != nil {
		errorMessage := rest_err.NewBadRequestError("Invalid user id")
		logger.Error(
			errorMessage.Message,
			err,
			journey,
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	err := uh.service.DeleteUser(uid)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

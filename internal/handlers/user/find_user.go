package user_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	user_response "github.com/rbaccaglini/simple_crud_golang/internal/models/response/user"
	"github.com/rbaccaglini/simple_crud_golang/internal/util/converter"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) FindAllUser(c *gin.Context) {
	journey := zap.String("journey", "FindAllUser")
	logger.Info("Init Find User", journey)

	ld, err := uh.service.FindAllUser()
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	listView := []user_response.UserResponse{}
	for _, d := range ld {
		viewUser := converter.ConvertDomainToResponse(d)
		listView = append(listView, viewUser)
	}

	logger.Info(fmt.Sprintf("There are %d users in DB", len(listView)), journey)

	c.JSON(http.StatusOK, listView)
}

func (uh *userHandlerInterface) FindUserById(c *gin.Context) {
	journey := zap.String("journey", "FindUserById")
	logger.Info("Init Find User", journey)

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

	d, err := uh.service.FindUserById(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(d))
}

func (uh *userHandlerInterface) FindUserByEmail(c *gin.Context) {
	journey := zap.String("journey", "FindUserByEmail")
	logger.Info("Init Find User controller", journey)
	userEmail := c.Param("email")

	d, err := uh.service.FindUserByEmail(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(d))
}

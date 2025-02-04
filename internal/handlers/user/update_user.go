package user_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/internal/models/domain"
	user_request "github.com/rbaccaglini/simple_crud_golang/internal/models/request/user"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	validation_err "github.com/rbaccaglini/simple_crud_golang/pkg/utils/validation/validator_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) UpdateUser(c *gin.Context) {
	journey := zap.String("journey", "UpdateUser")
	logger.Info("Init Update User", journey)

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

	var UserRequest user_request.UserUpdateRequest
	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("There are some incorrect fields", err, journey)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		errRest := validation_err.ValidateUserError(err)
		c.JSON(restErr.Code, errRest)
		return
	}

	domain := domain.NewUserUpdateDomain(
		UserRequest.Name,
		UserRequest.Age,
	)

	if err := uh.service.UpdateUser(domain, uid); err != nil {
		logger.Error("calling service", err, journey)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)

}

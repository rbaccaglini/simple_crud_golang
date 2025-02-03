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

	"go.uber.org/zap"
)

func (uh *userHandlerInterface) CreateUser(c *gin.Context) {

	var JOURNEY = zap.String("journey", "CreateUser")

	logger.Info("CreateUser function called", JOURNEY)

	var UserRequest user_request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		logger.Error("There are some incorrect fields", err, JOURNEY)
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect fields, error=%s", err.Error()))
		errRest := validation_err.ValidateUserError(err)
		c.JSON(restErr.Code, errRest)
		return
	}

	domain := domain.NewUserDomain(
		UserRequest.Email,
		UserRequest.Password,
		UserRequest.Name,
		UserRequest.Age,
	)

	d, err := uh.service.CreateUser(domain)
	if err != nil {
		logger.Error("calling service", err, JOURNEY)
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User created with success [id=%s]", d.GetID()), JOURNEY)
	c.JSON(http.StatusCreated, d.ConvertDomainToResponse())

}

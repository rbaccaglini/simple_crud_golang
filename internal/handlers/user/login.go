package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_request "github.com/rbaccaglini/simple_crud_golang/internal/models/request/user"
	"github.com/rbaccaglini/simple_crud_golang/internal/util/converter"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/logger"
	"github.com/rbaccaglini/simple_crud_golang/pkg/utils/rest_err"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) Login(c *gin.Context) {
	journey := zap.String("journey", "Login")
	logger.Info("Init login", journey)

	mLogin := &user_request.LoginRequest{}
	if err := c.ShouldBindJSON(mLogin); err != nil {
		logger.Error("some fields with error", err, journey)
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid json body"))
		return
	}

	t, ud, err := uh.service.Login(mLogin.Email, mLogin.Password)
	if err != nil {
		logger.Error("error on create token", err, journey)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Header("Authorization", t)
	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(ud))
}

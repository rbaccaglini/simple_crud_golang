package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/controller"
	"github.com/rbaccaglini/simple_crud_golang/src/model"
)

func InitRouter(r *gin.RouterGroup, controller controller.UserControllerInterface) {
	r.GET("/user/:userId", model.VerifyToken, controller.FindUserById)
	r.GET("/user/byEmail/:userEmail", model.VerifyToken, controller.FindUserByEmail)
	r.GET("/user", model.VerifyToken, controller.FindAllUsers)
	r.POST("/user", model.VerifyToken, controller.CreateUser)
	r.PUT("/user/:userId", model.VerifyToken, controller.UpdateUser)
	r.DELETE("/user/:userId", model.VerifyToken, controller.DeleteUser)

	r.POST("/login", controller.Login)
}

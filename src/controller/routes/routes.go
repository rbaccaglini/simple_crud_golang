package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/controller"
)

func InitRouter(r *gin.RouterGroup, controller controller.UserControllerInterface) {
	r.GET("/user/:userId", controller.FindUserById)
	r.GET("/user/byEmail/:userEmail", controller.FindUserByEmail)
	r.GET("/user", controller.FindAllUsers)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)

	r.POST("/login", controller.Login)
}

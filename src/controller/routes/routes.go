package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rbaccaglini/simple_crud_golang/src/controller"
)

func InitRouter(r *gin.RouterGroup, controller controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.GET("/users", controller.FindAllUsers)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}

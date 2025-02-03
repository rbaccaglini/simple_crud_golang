package routes

import (
	"github.com/gin-gonic/gin"
	user_handler "github.com/rbaccaglini/simple_crud_golang/internal/handlers/user"
)

func InitRouter(r *gin.RouterGroup, handler user_handler.UserHandlerInterface) {
	r.GET("/user", handler.FindAllUser)
	r.GET("/user/:userId", handler.FindUserById)
	r.GET("/user/email/:email", handler.FindUserByEmail)

	r.POST("/user", handler.CreateUser)
	r.DELETE("/user/:userId", handler.DeleteUser)
	r.PUT("/user/:userId", handler.UpdateUser)

	r.POST("/login", handler.Login)
}

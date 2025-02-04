package routes

import (
	"github.com/gin-gonic/gin"
	user_handler "github.com/rbaccaglini/simple_crud_golang/internal/handlers/user"
	"github.com/rbaccaglini/simple_crud_golang/internal/middleware"
)

func InitRouter(r *gin.RouterGroup, handler user_handler.UserHandlerInterface) {

	r.POST("/login", handler.Login)

	protected := r.Group("/user", middleware.VerifyToken)
	{
		protected.GET("", handler.FindAllUser)
		protected.GET("/:userId", handler.FindUserById)
		protected.GET("/email/:email", handler.FindUserByEmail)
		protected.POST("", handler.CreateUser)
		protected.DELETE("/:userId", handler.DeleteUser)
		protected.PUT("/:userId", handler.UpdateUser)
	}

}

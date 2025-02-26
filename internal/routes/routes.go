package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	user_handler "github.com/rbaccaglini/simple_crud_golang/internal/handlers/user"
)

func InitRouter(r *gin.RouterGroup, handler user_handler.UserHandlerInterface) {

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Permite qualquer origem
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // Desative credenciais para maior segurança
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", handler.Login)
	r.OPTIONS("/login", handleOptions)

	// protected := r.Group("/user", middleware.VerifyToken)
	protected := r.Group("/user")
	{
		protected.GET("", handler.FindAllUser)
		protected.GET("/:userId", handler.FindUserById)
		protected.GET("/email/:email", handler.FindUserByEmail)
		protected.POST("", handler.CreateUser)
		protected.DELETE("/:userId", handler.DeleteUser)
		protected.PUT("/:userId", handler.UpdateUser)

		protected.OPTIONS("", handleOptions)
		protected.OPTIONS("/:id", handleOptions)
		protected.OPTIONS("/email/:email", handleOptions)
	}

}

func handleOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	c.Status(204)
}

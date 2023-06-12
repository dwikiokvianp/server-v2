package routes

import (
	"github.com/gin-gonic/gin"
	"server-v2/controllers"
	"server-v2/utils"
)

func AuthRoutes(c *gin.Engine) {
	authGroup := c.Group("/auth")
	{
		authGroup.Use(utils.AuthMiddleware)
		authGroup.POST("/register", controllers.RegisterUser)
		authGroup.POST("/login", controllers.LoginUser)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"server-v2/controllers"
)

func UserRoutes(c *gin.Engine) {
	userGroup := c.Group("/user")
	{
		userGroup.GET("/", controllers.GetAllUser)
	}
}

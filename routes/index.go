package routes

import "github.com/gin-gonic/gin"

func Routes(c *gin.Engine) {
	AuthRoutes(c)
	TransactionRoutes(c)
}

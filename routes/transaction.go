package routes

import (
	"github.com/gin-gonic/gin"
	"server-v2/controllers"
	"server-v2/utils"
)

func TransactionRoutes(c *gin.Engine) {
	transactionGroup := c.Group("/transaction")
	{
		transactionGroup.Use(utils.AuthMiddleware)
		transactionGroup.GET("/", controllers.GetAllTransactions)
		transactionGroup.POST("/", controllers.CreateTransactions)
		transactionGroup.GET("/:id", controllers.GetByIdTransaction)
	}
}

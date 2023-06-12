package controllers

import (
	"github.com/gin-gonic/gin"
	"server-v2/config"
	"server-v2/models"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction

	config.DB.Find(&transactions)

	c.JSON(200, gin.H{
		"data": transactions,
	})
}

func GetByIdTransaction(c *gin.Context) {
	var transaction models.Transaction
	id := c.Param("id")
	config.DB.Find(&transaction, id)

	c.JSON(200, gin.H{
		"data": transaction,
	})
}

func CreateTransactions(c *gin.Context) {
	var transaction models.Transaction
	err := c.BindJSON(&transaction)
	if err != nil {
		return
	}

	config.DB.Create(&transaction)

	c.JSON(200, gin.H{
		"data": transaction,
	})
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"server-v2/config"
	"server-v2/models"
)

func GetAllUser(c *gin.Context) {
	var userList []models.User

	queryParams := c.Request.URL.Query()
	role := queryParams.Get("role")
	if role != "" {
		err := config.DB.Where("role_id = ?", role).Preload("Role").Preload("Detail").Find(&userList).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, gin.H{
		"data": userList,
	})
}

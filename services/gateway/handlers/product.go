package handlers

import "github.com/gin-gonic/gin"

func GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetProducts called",
	})
}

func CreateProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateProduct called",
	})
}

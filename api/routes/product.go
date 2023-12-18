package routes

import (
	"github.com/alfatahh54/create-transaction/model"
	"github.com/alfatahh54/create-transaction/services"
	"github.com/gin-gonic/gin"
)

func init() {
	go MainRoute.NewRoute("GET", "/product", GetAllProduct)
}

func GetAllProduct(c *gin.Context) {
	var products []model.Product
	err := services.GetAllModel(&products)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success get data products",
		"data":    products,
	})
}

package routes

import (
	"strconv"
	"time"

	"github.com/alfatahh54/create-transaction/api/model"
	"github.com/alfatahh54/create-transaction/api/services"
	"github.com/gin-gonic/gin"
)

func init() {
	go MainRoute.NewRoute("POST", "/transaction", CreateTransaction)
}

func CreateTransaction(c *gin.Context) {
	var body model.CreateTransactionBody

	err := c.Bind(&body)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}
	ids := make([]uint, 0)
	for _, p := range body.ProductList {
		ids = append(ids, p.ProductID)
	}
	var products []model.Product
	err = services.GetModelByIds(&products, ids)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}
	var transactions []model.Transaction
	transactionCode := time.Now().Unix()
	for _, i := range body.ProductList {
		var transaction model.Transaction
		for _, p := range products {
			if i.ProductID == p.ID {
				transaction.Qty = i.Qty
				transaction.Total = i.Qty * p.Price
				transaction.ProductID = p.ID
				transaction.Price = p.Price
			}
		}
		transaction.CustomerName = body.CustomerName
		transaction.TransactionCode = strconv.Itoa(int(transactionCode))
		transactions = append(transactions, transaction)
	}
	err = services.Create(&transactions)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "Server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success create data business",
		"data":    strconv.Itoa(int(transactionCode)),
	})
}

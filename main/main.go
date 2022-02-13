package main

import (
	"gostockapi.com/stock"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/stocks", stock.GetStocks)
	router.GET("/stocks/:id", stock.GetStockById)
	router.POST("/stocks", stock.AddStock)
	router.DELETE("/stocks/:id", stock.DeleteStockById)
	router.POST("/stocks/:id", stock.UpdateStock)

	router.Run("localhost:7000")

}

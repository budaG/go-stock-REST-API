package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type stock struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
}

func main() {
	router := gin.Default()
	router.GET("/stocks", getStocks)
	router.GET("/stocks/:id", getStockById)
	router.POST("/stocks", addStock)
	router.DELETE("/stocks/:id", deleteStockById)
	router.POST("/stocks/:id", updateStock)

	router.Run("localhost:7000")
}

var stocks = []stock{
	{ID: 1, Name: "Boeing", Symbol: "BA", Open: 218.01, High: 221.41, Low: 210.27},
	{ID: 2, Name: "Delta Air Lines", Symbol: "DAL", Open: 43.93, High: 44.80, Low: 41.92},
	{ID: 3, Name: "Uber", Symbol: "UBER", Open: 37.98, High: 38.64, Low: 34.48},
	{ID: 4, Name: "Sundial Growers", Symbol: "SNDL", Open: 0.56, High: 0.66, Low: 0.55},
	{ID: 5, Name: "Tesla", Symbol: "TSLA", Open: 909.39, High: 915.96, Low: 850.71},
}

func getStocks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stocks)
}

func addStock(c *gin.Context) {
	var newStock stock

	if err := c.BindJSON(&newStock); err != nil {
		return
	}

	// add stock to the slice.
	stocks = append(stocks, newStock)
	c.IndentedJSON(http.StatusCreated, newStock)
}

func getStockById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "bad id"})
		return
	}

	for _, stock := range stocks {
		if stock.ID == id {
			c.IndentedJSON(http.StatusOK, stock)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "stock not found"})
}

func deleteStockById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "bad id"})
		return
	}

	exist, resStocks := deleteStock(c, id)
	stocks = resStocks

	if exist {
		c.IndentedJSON(http.StatusOK, stocks)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message:": "stock not found"})

	}

}

func deleteStock(c *gin.Context, id int) (bool, []stock) {
	resStocks := []stock{}
	exist := false
	for _, stock := range stocks {
		if stock.ID == id {
			exist = true
		} else {
			resStocks = append(resStocks, stock)
		}
	}
	return exist, resStocks
}

func updateStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: ": "bad id"})
		return
	}

	var reqStock stock

	if err := c.BindJSON(&reqStock); err != nil {
		return
	}

	exist, resStocks := deleteStock(c, id)

	if exist {
		resStocks = append(resStocks, reqStock)
		stocks = resStocks
		c.IndentedJSON(http.StatusOK, stocks)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message:": "stock not found"})

	}
}

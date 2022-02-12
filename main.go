package main

import (
	"net/http"

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

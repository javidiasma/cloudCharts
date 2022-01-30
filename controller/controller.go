package controller

import (
	"cloudCharts/DAL"
	"cloudCharts/csvReader"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateDB(c *gin.Context) {
	fmt.Println(1)
	_, err := csvReader.CsvReader()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "done"})
		return
	}
}
func GetDB(c *gin.Context) {
	fmt.Println(2)
	data, err := DAL.GetDatabase()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, data)
		return
	}
}

func First(c *gin.Context) { // compare between two games
	firstGame := c.Param("first")
	secondGame := c.Param("second")
	firstGameSales := DAL.GenerateBarItemsByName(firstGame)
	secondGameSales := DAL.GenerateBarItemsByName(secondGame)
	chart := DAL.First(firstGame, secondGame, firstGameSales, secondGameSales)
	c.HTML(http.StatusOK, chart, nil)
}

func Second(c *gin.Context) {

}
func Third(c *gin.Context) {

}
func Forth(c *gin.Context) {

}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type spice struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Level      float32 `json:"level"`
	Substitute string  `json:"substitute"`
}

//seed data
var spices = []spice{
	{Id: 1, Name: "Oregano", Level: 50, Substitute: "Rosemary"},
	{Id: 2, Name: "Rosemary", Level: 25, Substitute: "Oregano"},
	{Id: 3, Name: "Sage", Level: 75, Substitute: "Thyme"},
	{Id: 4, Name: "Thyme", Level: 125, Substitute: "Sage"},
}

func main() {
	router := gin.Default()

	router.GET("/spices", getSpices)
	router.POST("/spices", postSpice)

	router.Run("localhost:8080")
}

func getSpices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, spices)
}

func postSpice(c *gin.Context) {
	var newSpice spice

	if err := c.BindJSON(&newSpice); err != nil {
		fmt.Printf("Error processing request: %v", err.Error())
		return
	}

	newSpice.Id = getNextId()

	spices = append(spices, newSpice)
	c.IndentedJSON(http.StatusCreated, newSpice)
}

func getNextId() int {
	maxId := 0
	for _, spiceEntry := range spices {
		if spiceEntry.Id > maxId {
			maxId = spiceEntry.Id
		}
	}

	return maxId + 1
}

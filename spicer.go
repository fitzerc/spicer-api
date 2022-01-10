package main

import (
	"fmt"
	"net/http"
	"spicer-api/shared"

	"github.com/gin-gonic/gin"
)

var dal shared.SpiceDal

func main() {
	dal = shared.FsSpiceDal{}
	router := gin.Default()

	router.GET("/spices", getSpices)
	router.POST("/spices", postSpice)

	router.Run("localhost:8080")
}

func getSpices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dal.ReadSpices())
}

func postSpice(c *gin.Context) {
	var newSpice shared.Spice

	if err := c.BindJSON(&newSpice); err != nil {
		fmt.Printf("Error processing request: %v", err.Error())
		return
	}

	newSpice.Id = getNextId(dal.ReadSpices())

	dal.WriteSpice(newSpice)
	c.IndentedJSON(http.StatusCreated, newSpice)
}

func getNextId(s []shared.Spice) int {
	maxId := 0
	for _, spiceEntry := range s {
		if spiceEntry.Id > maxId {
			maxId = spiceEntry.Id
		}
	}

	return maxId + 1
}

package main

import (
	"bookshelf/controllers"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

func main() {
	router := gin.Default()
	router.Get("/:id", func(c *gin.Context) {
		// process param
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			c.JSON(400, err)
			return
		}
		if id <= 0 {
			c.JSON(400, gin.H{"status": "id should be bigger than 0"})
			return
		}
		// process data
		ctrl := controllers.NewUser()
		result := ctrl.Get(id)
		if result == nil || reflect.ValueOf(result).IsNil() {
			c.JSON(404, gin.H{})
			return
		}

		c.JSON(200, result)
	})
	router.Run(":8080")
}

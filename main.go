package main

import (
	"github.com/gin-gonic/gin"
	"github.com/watsonso/bookshelf/controllers"
	"net/http"
	_"reflect"
	_"strconv"
	"fmt"
)

func main() {
	router := gin.Default()
	fmt.Println("check 1")
	router.LoadHTMLGlob("views/*")
	fmt.Println("check 2")
	router.GET("/", func(c *gin.Context) {
		ctrl := controllers.NewTask()
		fmt.Println("check get")

		tasks := ctrl.GetAll()

		fmt.Println("check get2", tasks)
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"task": tasks,
		})
	})

	router.POST("/", func(c *gin.Context) {
		text := c.PostForm("text")

		ctrl := controllers.NewTask()

		fmt.Println("check text", text)
		//if text == '' {
		//
		//}
		ctrl.Create(text)

		tasks := ctrl.GetAll()

		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"tasks": tasks,
		})
	})

	//router.GET("/:id", func(c *gin.Context) {
	//	// process param
	//	n := c.Param("id")
	//	id, err := strconv.Atoi(n)
	//	if err != nil {
	//		c.JSON(400, err)
	//		return
	//	}
	//	if id <= 0 {
	//		c.JSON(400, gin.H{"status": "id should be bigger than 0"})
	//		return
	//	}
	//	// process data
	//	ctrl := controllers.NewUser()
	//	result := ctrl.Get(id)
	//	if result == nil || reflect.ValueOf(result).IsNil() {
	//		c.JSON(404, gin.H{})
	//		return
	//	}

	//	c.JSON(200, result)
	//})
	fmt.Println("check 3")
	router.Run(":8080")
}

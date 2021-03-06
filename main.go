package main

import (
	"github.com/gin-gonic/gin"
	"github.com/watsonso/bookshelf/controllers"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/post", func(c *gin.Context) {
		text := c.PostForm("text")

		ctrl := controllers.NewTask()

		if text == "" {
			c.Redirect(http.StatusMovedPermanently, "/")
			c.Abort()
		}
		ctrl.Create(text)

		c.Redirect(http.StatusMovedPermanently, "/")

	})

	router.POST("/delete", func(c *gin.Context) {
		stringID := c.PostForm("id")

		ctrl := controllers.NewTask()
		id, _ := strconv.Atoi(stringID)
		ctrl.Delete(id)

		tasks := ctrl.GetAll()

		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"tasks": tasks,
		})

	})

	router.GET("/", func(c *gin.Context) {
		ctrl := controllers.NewTask()

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
	router.Run(":8080")
}

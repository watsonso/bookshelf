package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 1. http://localhost:8080/ へアクセスすると「Hello world」と表示する。
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	//2. http://localhost:8080/hoge へアクセスすると、「fuga」と表示する。
	r.GET("/hoge", func(c *gin.Context) {
		c.String(200, "fuga")
	})
	r.Run(":8080")
}

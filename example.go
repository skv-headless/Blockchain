package main

import "github.com/gin-gonic/gin"
import "fmt"

var i int = 2

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		i = i + 1
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("i = %d", i),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

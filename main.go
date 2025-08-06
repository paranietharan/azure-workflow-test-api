package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()
	r.GET("/test", TestEndPoint)
	r.GET("/", RootEndPoint)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func RootEndPoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "Hello World",
		"endPoint": "/",
	})
}

func TestEndPoint(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"Message":  "Hello World - TestEndPoint",
			"endPoint": "/test",
		},
	)
}

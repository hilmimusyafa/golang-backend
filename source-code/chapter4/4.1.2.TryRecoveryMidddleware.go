package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Using gin.Default() which includes Recovery middleware
	r := gin.Default()

	r.GET("/safe", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a safe route"})
	})

	r.GET("/cause-panic", func(c *gin.Context) {
		// Simulating panic: attempting to access index outside slice bounds
		var data []string
		_ = data[0] // This will cause panic!
		c.JSON(http.StatusOK, gin.H{"message": "This line will never be reached"})
	})

	r.Run(":8080")
}
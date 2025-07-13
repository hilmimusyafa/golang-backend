package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Includes Recovery middleware by default

	r.GET("/cause-panic", func(c *gin.Context) {
		// Simulate a panic (e.g., trying to access nil pointer)
		var data []int
		_ = data[0] // This will cause a panic
		c.JSON(http.StatusOK, gin.H{"message": "This will not be reached"})
	})

	r.Run(":8080")
}
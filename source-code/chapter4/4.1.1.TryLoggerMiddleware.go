package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a Gin machine without built-in middleware
	r := gin.New()

	// Add Logger middleware manually
	r.Use(gin.Logger())

	// Simple route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
	})

	r.Run(":8080")
}
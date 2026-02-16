package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Endpoint with custom header Authorization
	r.GET("/protected", func(c *gin.Context) {
		
		// Get header Authorization from request
		authToken := c.GetHeader("Authorization")
		
		// Simple token validation (on production use JWT)
		expectedToken := "Bearer secret123"
		
		if authToken != expectedToken {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Invalid token",
			})
			return
		}

		// Set custom response header
		c.Header("X-Request-ID", "req-12345")
		c.Header("X-Api-Version", "v1.0")
		
		c.JSON(http.StatusOK, gin.H{
			"message": "Access granted",
			"data": gin.H{
				"user": "Hilmi",
				"role": "admin",
			},
		})
	})

	r.Run(":8080")
}
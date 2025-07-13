package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth middleware (revisited)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Unauthorized access",
				"code":    http.StatusUnauthorized,
			})
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.GET("/protected", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "You accessed protected data!"})
	})

	r.Run(":8080")
}
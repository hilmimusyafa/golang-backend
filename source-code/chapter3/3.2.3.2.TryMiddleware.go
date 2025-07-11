package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Simple logging middleware: records every incoming request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()           // Record start time
		c.Next()                      // Call the next handler
		duration := time.Since(start) // Calculate duration
		fmt.Printf("[%s] %s %s %v\n", c.Request.Method, c.Request.URL.Path, c.ClientIP(), duration)
	}
}

// Simple authentication middleware: checks the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			// If the token is invalid, stop the request and send a 401 status
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		c.Next() // Token is valid, proceed to the next handler
	}
}           

// Example of using middleware in Gin
func main() {
	r := gin.Default()

	// 1. Global Middleware : applies to all routes
	r.Use(LoggerMiddleware())

	// 2. Public route without authentication
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a public area"})
	})

	// 3. Group Middleware : only for routes within the /private group
	private := r.Group("/private")
	private.Use(AuthMiddleware())
	{
		private.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "This is secret data"})
		})
		private.POST("/settings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Settings updated successfully"})
		})
	}

	// 4. Route Middleware : only for one route
	r.GET("/admin-dashboard", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the admin dashboard"})
	})

	r.Run(":8080")
}

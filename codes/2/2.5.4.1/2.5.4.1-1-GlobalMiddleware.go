package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware for logging every request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Note the start time of the request
		startTime := time.Now()

		// Note the request information
		method := c.Request.Method
		path := c.Request.URL.Path

		fmt.Printf("[%s] Request: %s %s\n", 
			startTime.Format("2006-01-02 15:04:05"), 
			method, 
			path,
		)
		
		// Call the next handler in the chain
		c.Next()

		// The code below is executed AFTER the handler completes
		// Calculate the duration and log the response status
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		fmt.Printf("[%s] Response: %s %s - Status: %d - Duration: %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			statusCode,
			duration,
		)
	}
}

// Middleware for recovery from panic
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("PANIC RECOVERED: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
			}
		}()

		c.Next()
	}
}

func main() {
	// Use gin.New() for a bare router with no built-in middleware
	r := gin.New()

	// Register global middleware
	// The order is important: the middleware will be executed in the order they are registered
	r.Use(RecoveryMiddleware())
	r.Use(LoggerMiddleware())

	// Normal route
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "List of users",
			"data": []string{"Alice", "Bob", "Charlie"},
		})
	})

	// Route that will panic (for test recovery)
	r.GET("/panic", func(c *gin.Context) {
		panic("Something went wrong!")
	})

	r.Run(":8080")
}
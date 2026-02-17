package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIDMiddleware adds a unique ID to each request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

// LoggerMiddleware records request details in a structured format
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		requestID, _ := c.Get("request_id")

		c.Next()

		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		fmt.Printf("[%s] [%s] %s %s - Status: %d - Duration: %v\n",
			requestID,
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			statusCode,
			duration,
		)
	}
}

// ErrorHandlerMiddleware handles errors centrally
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Cek apakah ada error yang di-set di context
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			
			// Log error
			fmt.Printf("ERROR: %v\n", err.Err)

			// Kirim response error yang konsisten
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Internal server error",
				"message": err.Err.Error(),
			})
		}
	}
}

// RateLimitMiddleware limits the number of requests per IP
func RateLimitMiddleware(maxRequests int, duration time.Duration) gin.HandlerFunc {
	type client struct {
		count     int
		lastReset time.Time
	}

	clients := make(map[string]*client)

	return func(c *gin.Context) {
		ip := c.ClientIP()

		if _, exists := clients[ip]; !exists {
			clients[ip] = &client{
				count:     0,
				lastReset: time.Now(),
			}
		}

		clientInfo := clients[ip]

		// Reset counter jika sudah melewati durasi
		if time.Since(clientInfo.lastReset) > duration {
			clientInfo.count = 0
			clientInfo.lastReset = time.Now()
		}

		// Cek apakah sudah melewati limit
		if clientInfo.count >= maxRequests {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Please try again later.",
			})
			c.Abort()
			return
		}

		clientInfo.count++
		c.Next()
	}
}

// TimeoutMiddleware provides a timeout to the handler
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		done := make(chan bool, 1)

		go func() {
			c.Next()
			done <- true
		}()

		select {
		case <-done:
			// Handler selesai tepat waktu
			return
		case <-time.After(timeout):
			// Timeout terjadi
			c.JSON(http.StatusRequestTimeout, gin.H{
				"error": "Request timeout",
			})
			c.Abort()
		}
	}
}

func main() {
	r := gin.New()

	// Use production-ready middleware
	r.Use(RequestIDMiddleware())
	r.Use(LoggerMiddleware())
	r.Use(ErrorHandlerMiddleware())
	r.Use(RateLimitMiddleware(10, time.Minute)) // 10 req/menit per IP

	// Routes
	r.GET("/api/users", TimeoutMiddleware(5*time.Second), func(c *gin.Context) {
		// Slow process simulation
		time.Sleep(2 * time.Second)
		
		c.JSON(200, gin.H{
			"message": "Success",
		})
	})

	r.Run(":8080")
}
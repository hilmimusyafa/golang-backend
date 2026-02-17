package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware for simple rate limiting
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simulasi rate limiting
		// Di production, gunakan Redis atau library seperti golang.org/x/time/rate
		time.Sleep(100 * time.Millisecond) // Simulasi delay

		c.Next()
	}
}

// Middleware for API versioning
func APIVersionMiddleware(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-API-Version", version)
		c.Next()
	}
}

// Simple middleware for authentication
func SimpleAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer valid-token" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Public route without additional middleware
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to API",
		})
	})

	// Group API v1 with middleware rate limiting and versioning
	v1 := r.Group("/api/v1")
	v1.Use(RateLimitMiddleware())
	v1.Use(APIVersionMiddleware("1.0"))
	{
		// Endpoint publik dalam v1
		v1.GET("/products", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "List of products",
				"version": "v1",
				"data": []string{"Product A", "Product B"},
			})
		})

		// Sub-group untuk endpoint yang butuh autentikasi
		authenticated := v1.Group("/")
		authenticated.Use(SimpleAuthMiddleware())
		{
			authenticated.GET("/orders", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Your orders",
					"data": []string{"Order #1", "Order #2"},
				})
			})

			authenticated.POST("/orders", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"message": "Order created",
				})
			})
		}
	}

	// Group API v2 with different middleware
	v2 := r.Group("/api/v2")
	v2.Use(APIVersionMiddleware("2.0"))
	{
		v2.GET("/products", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "List of products",
				"version": "v2",
				"data": []gin.H{
					{"id": 1, "name": "Product A", "price": 100},
					{"id": 2, "name": "Product B", "price": 200},
				},
			})
		})
	}

	// Group admin with multiple middleware
	admin := r.Group("/admin")
	admin.Use(SimpleAuthMiddleware())
	admin.Use(func(c *gin.Context) {
		// Middleware khusus admin
		c.Header("X-Admin-Panel", "true")
		c.Next()
	})
	{
		admin.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "List of all users",
				"data": []string{"User 1", "User 2", "User 3"},
			})
		})

		admin.DELETE("/users/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted",
				"user_id": userID,
			})
		})
	}

	r.Run(":8080")
}
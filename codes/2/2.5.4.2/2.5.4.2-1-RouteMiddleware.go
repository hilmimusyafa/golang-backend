package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//Middleware for simple authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Take header Authorization
		authHeader := c.GetHeader("Authorization")

		// Format validation: "Bearer <token>"
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort() // Stop subsequent middleware/handler execution
			return
		}

		// Parse token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization format. Use: Bearer <token>",
			})
			c.Abort()
			return
		}

		token := parts[1]

		// Vsimple token validation (in production use JWT)
		if token != "secret-token-123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		// Save user information to context for handler use
		c.Set("user_id", 1)
		c.Set("username", "hilmi")

		// Proceed to the handler
		c.Next()
	}
}

// Middleware to check the admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the user_id that has been stored by AuthMiddleware
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not authenticated",
			})
			c.Abort()
			return
		}

		// Role checking simulation (in production, take from database)
		// Assume the user with ID 1 is an admin
		if userID.(int) != 1 {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Admin access required",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Public route (without middleware)
	r.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "This is a public endpoint",
		})
	})

	// Routes that require authentication (use AuthMiddleware)
	r.GET("/profile", AuthMiddleware(), func(c *gin.Context) {
		// Ambil data user dari context
		username, _ := c.Get("username")
		
		c.JSON(http.StatusOK, gin.H{
			"message": "Your profile",
			"user":    username,
		})
	})

	// Routes that require authentication AND an admin role (chain middleware)
	r.GET("/admin/dashboard", AuthMiddleware(), AdminMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to admin dashboard",
			"data": gin.H{
				"total_users":  100,
				"total_orders": 250,
			},
		})
	})

	// Route to delete (requires admin)
	r.DELETE("/users/:id", AuthMiddleware(), AdminMiddleware(), func(c *gin.Context) {
		userID := c.Param("id")
		
		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted successfully",
			"user_id": userID,
		})
	})

	r.Run(":8080")
}
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define struct for user registration with validation tags
type RegisterUser struct {
	Username string `json:"username" binding:"required,min=5,max=20"` // Must be 5-20 characters
	Email    string `json:"email" binding:"required,email"`         // Must be a valid email format
	Password string `json:"password" binding:"required,min=6"`      // Must be at least 6 characters
	Age      int    `json:"age" binding:"gte=18"`                     // Must be 18 or older
}

func main() {
	r := gin.Default()

	// Endpoint for user registration with validation
	r.POST("/register", func(c *gin.Context) {
		var user RegisterUser
		if err := c.ShouldBindJSON(&user); err != nil {
			// If validation fails, Gin's binding error will contain details
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If validation is successful, proceed with user creation
		c.JSON(http.StatusCreated, gin.H{
			"message":  "User registered successfully!",
			"username": user.Username,
			"email":    user.Email,
			"age":      user.Age,
		})
	})

	r.Run(":8080")
}
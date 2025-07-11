package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct to map incoming JSON data
type User struct {
	ID       string `json:"id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
}

func main() {
	r := gin.Default()

	// Endpoint to create a new user from JSON data
	r.POST("/users", func(c *gin.Context) {
		var user User
		// Bind JSON data from request body to the 'user' struct
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If binding is successful, process the data
		c.JSON(http.StatusCreated, gin.H{
			"message":  "User created successfully",
			"user_id":   user.ID,
			"username": user.Username,
			"email":    user.Email,
		})
	})

	r.Run(":8080")
}
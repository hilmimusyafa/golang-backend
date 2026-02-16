package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister defines the data structure expected from the client.
type UserRegister struct {
	// binding:"required" ensure this field cannot be empty
	Username string `json:"username" binding:"required"`
	
	// binding:"email" validate email format
	Email    string `json:"email" binding:"required,email"`
	
	// binding:"min=8" ensure password has at least 8 characters
	Password string `json:"password" binding:"required,min=8"`
	
	// Field optional without validation
	Age      int    `json:"age"` 
}

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var input UserRegister

		// 1. Binding & Validation
		// ShouldBindJSON will automatically bind the JSON body to the 'input' struct and perform validation based on the tags
		if err := c.ShouldBindJSON(&input); err != nil {
			// If validation fails, return a 400 Bad Request with the error message
			// gin.H{"error": err.Error()} will create a JSON response with the error details
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 2. Processing Data Simulation (example: save to database)
		// In here, input.Username dll will safely used
		c.JSON(http.StatusOK, gin.H{
			"message":  "User berhasil didaftarkan",
			"username": input.Username,
			"email":    input.Email,
		})
	})

	// Binding Query String example (etc: /search?q=golang&page=1)
	r.GET("/search", func(c *gin.Context) {
		var query struct {
			Q    string `form:"q" binding:"required"` // Gunakan tag 'form' untuk query param
			Page int    `form:"page"`
		}

		// ShouldBindQuery specifically for retrieving parameters from URLs
		if err := c.ShouldBindQuery(&query); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"search_query": query.Q,
			"page":         query.Page,
		})
	})

	r.Run(":8080")
}
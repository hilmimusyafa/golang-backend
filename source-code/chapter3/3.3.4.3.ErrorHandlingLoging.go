package main

import (
	"log" // For simple logging
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Includes Logger and Recovery middleware

	r.POST("/process-data", func(c *gin.Context) {
		var input struct {
			Value int `json:"value"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			log.Printf("ERROR: Invalid JSON input: %v", err) // Log the error
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
			return
		}

		if input.Value < 0 {
			// Log a specific business logic error
			log.Printf("ERROR: Negative value received for processing: %d", input.Value)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Value cannot be negative"})
			return
		}

		// Simulate some processing that might fail
		if input.Value == 999 {
			log.Println("CRITICAL ERROR: Simulated database write failure!") // Log critical error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process data due to internal issue"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data processed successfully", "value": input.Value})
	})

	r.Run(":8080")
}
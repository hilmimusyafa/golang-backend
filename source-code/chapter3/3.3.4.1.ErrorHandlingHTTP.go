package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Simulate a database lookup error or user not found
		if id == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "User not found",
				"code":    http.StatusNotFound,
			})
			return // Stop processing
		}

		// Simulate an internal server error (e.g., database connection issue)
		if id == "server-error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "error",
				"message": "Internal server error occurred",
				"code":    http.StatusInternalServerError,
			})
			return // Stop processing
		}

		// If no error, return success
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "User data for ID: " + id,
			"code":    http.StatusOK,
		})
	})

	r.Run(":8080")
}
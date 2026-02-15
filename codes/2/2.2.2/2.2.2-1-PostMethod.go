package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Database simulation (in-memory)
var books = []Book{
	{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
	{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
}

func main() {
	r := gin.Default()

	// 2. POST: Adding new data
	r.POST("/books", func(c *gin.Context) {

		var newBook Book

		// Bind JSON from body request
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		// Generate simple ID (auto increment)
		newBook.ID = len(books) + 1

		// Add to slice (database insert )
		books = append(books, newBook)

		c.JSON(http.StatusCreated, gin.H{
			"message": "Book created successfully",
			"data":    newBook,
		})
	})

	r.Run(":8080")
}

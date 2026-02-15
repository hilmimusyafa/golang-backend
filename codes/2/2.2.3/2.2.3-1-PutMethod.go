package main

import (
	"net/http"
	"strconv"

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

	// 3. PUT: Change all data
	r.PUT("/books/:id", func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid book ID",
			})
			return
		}

		var updatedBook Book

		// Bind JSON from body request
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		// Call update function
		book, found := updateBookByID(id, updatedBook)

		if !found {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book updated successfully",
			"data":    book,
		})
	})

	r.Run(":8080")
}

// Function to update book by ID
func updateBookByID(id int, updatedData Book) (Book, bool) {

	for i, book := range books {
		if book.ID == id {
			updatedData.ID = id // Pastikan ID tetap sama
			books[i] = updatedData
			return updatedData, true
		}
	}

	return Book{}, false
}

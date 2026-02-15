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

	// 5. DELETE: Delete data
	r.DELETE("/books/:id", func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid book ID",
			})
			return
		}

		// Call delete function
		deletedBook, found := deleteBookByID(id)

		if !found {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book deleted successfully",
			"data":    deletedBook,
		})
	})

	r.Run(":8080")
}

// Function for delete book by ID
func deleteBookByID(id int) (Book, bool) {
	for i, book := range books {
		if book.ID == id {
			// Save deleted data for response 
			deletedBook := book

			// Delete from slice
			books = append(books[:i], books[i+1:]...)

			return deletedBook, true
		}
	}
	return Book{}, false
}
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

// Simulasi database (in-memory)
var books = []Book{
	{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
	{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
}

// Struct khusus untuk PATCH (pakai pointer agar bisa deteksi field yang dikirim)
type UpdateBookInput struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
}

func main() {
	r := gin.Default()

	// 4. PATCH: Update only piece of data
	r.PATCH("/books/:id", func(c *gin.Context) {

		idParam := c.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid book ID",
			})
			return
		}

		var input UpdateBookInput

		// Bind JSON parsial
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON data",
			})
			return
		}

		// Call function for partial update
		book, found := patchBookByID(id, input)

		if !found {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Book not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Book partially updated",
			"data":    book,
		})
	})

	r.Run(":8080")
}

// Function for partial update book by ID
func patchBookByID(id int, input UpdateBookInput) (Book, bool) {

	for i, book := range books {

		if book.ID == id {

			// Update only field sended by client
			if input.Title != nil {
				books[i].Title = *input.Title
			}

			if input.Author != nil {
				books[i].Author = *input.Author
			}

			return books[i], true
		}
	}

	return Book{}, false
}

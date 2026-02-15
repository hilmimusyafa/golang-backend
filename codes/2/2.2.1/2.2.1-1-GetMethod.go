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

func main() {
	r := gin.Default()

	// 1. GET: Take data
	r.GET("/books", func(c *gin.Context) {

		// Take data from database (simulated)
		bookData := getBookData()

		c.JSON(http.StatusOK, gin.H{
			"message": "Showing list of books",
			"data":    bookData,
		})
	})

	r.Run(":8080")
}


// Simulate fetching data from a database
func getBookData() []Book {

	books := []Book{
		{ID: 1, Title: "Clean Code", Author: "Robert C. Martin"},
		{ID: 2, Title: "Refactoring", Author: "Martin Fowler"},
		{ID: 3, Title: "Domain-Driven Design", Author: "Eric Evans"},
	}

	return books
}
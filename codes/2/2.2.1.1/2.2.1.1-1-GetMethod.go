package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 1. GET: Mengambil data
	r.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Showing list of books",
		})
	})

	r.Run(":8080") // running the server on port 8080
}
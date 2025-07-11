package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Endpoint to update product data based on ID
	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		// Get new product name data from form or JSON (simple)
		newName := c.DefaultPostForm("name", "Produk Tanpa Nama")
		c.JSON(http.StatusOK, gin.H{
			"message":  "Produk dengan ID " + id + " berhasil diperbarui.",
			"new_name": newName,
		})
	})

	router.Run(":8080")
}

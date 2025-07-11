package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Defines a route that accepts query parameters
    r.GET("/products", func(c *gin.Context) {
        category := c.Query("category") // Get the value of the query parameter "category"
        sort := c.Query("sort") // Get the value of the query parameter "sort"

        if category != "" && sort != "" {
            c.JSON(http.StatusOK, gin.H{
                "message": "Mendapatkan produk kategori: " + category + " dengan sort: " + sort,
            })
        } else if category != "" {
            c.JSON(http.StatusOK, gin.H{
                "message": "Mendapatkan produk kategori: " + category,
            })
        } else if sort != "" {
            c.JSON(http.StatusOK, gin.H{
                "message": "Mendapatkan produk dengan sort: " + sort,
            })
        } else {
            c.JSON(http.StatusOK, gin.H{
                "message": "Mendapatkan semua produk",
            })
        }
    })

    r.Run(":8080")
}
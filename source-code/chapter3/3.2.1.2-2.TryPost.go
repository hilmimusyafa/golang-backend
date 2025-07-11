package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    // Endpoint for making new users
    router.POST("/create-user", func(c *gin.Context) {
        // Example : receive username data from form or JSON (simplified)
        name := c.DefaultPostForm("name", "Anonymous")
        c.JSON(http.StatusCreated, gin.H{
            "message": "User created successfully.",
            "user":    name,
        })
    })

    // Endpoint for making new products
    router.POST("/create-product", func(c *gin.Context) {
        // Example : receive product name data from form or JSON (simplified)
        product := c.DefaultPostForm("product", "Unknown Product")
        c.JSON(http.StatusCreated, gin.H{
            "message": "Product created successfully.",
            "product": product,
        })
    })
    
    router.Run(":8080")
}
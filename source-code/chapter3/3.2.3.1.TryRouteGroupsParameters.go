package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Create a group for admin routes
    admin := r.Group("/admin")
    {
        admin.GET("/users", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "List of admin users"})
        })
        admin.POST("/products", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "New products added by admin"})
        })
    }

    // Creating a group for API version 1
    apiV1 := r.Group("/api/v1")
    {
        apiV1.GET("/data", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "Data from API v1"})
        })
    }

    // Routes outside the group can still be created as usual.
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Home page"})
    })

    r.Run(":8080")
}
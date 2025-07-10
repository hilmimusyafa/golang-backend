package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Define a route with dynamic parameters :id
    router.GET("/users/:id", func(c *gin.Context) {
        // Get the value of the 'id' parameter from the URL
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
            "message": "Getting user data with ID: " + id,
        })
    })
    router.Run(":8080")
}
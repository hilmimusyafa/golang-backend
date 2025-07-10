package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin with default middleware (Logger and Recovery)
    r := gin.Default()

    // Define route for root path ("/") with GET method
    r.GET("/", func(c *gin.Context) {
        // Send JSON response with status 200 OK
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello From Gin!",
        })
    })

    // Run server on port 8080
    r.Run(":8080")
}
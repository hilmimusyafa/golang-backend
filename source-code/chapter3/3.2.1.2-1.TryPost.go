package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Define a route for the POST method in the path "/create-user"
    router.POST("/create-user", func(c *gin.Context) {
        // ... Logic to create a new user will go here ...
        // For now, we just send a confirmation response
        c.JSON(http.StatusCreated, gin.H{
            "message": "User created successfully.",
        })
    })
    router.Run(":8080")
}
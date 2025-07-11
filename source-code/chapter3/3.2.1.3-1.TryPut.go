package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Defines a route for the PUT method at the path "/users/:id"
    // :id is a dynamic parameter that can be taken from the URL
    router.PUT("/users/:id", func(c *gin.Context) {
        id := c.Param("id") // Retrieving the "id" parameter value from the URL
        // The logic for updating a user with a specific ID will be here.
        c.JSON(http.StatusOK, gin.H{
            "message": "User with ID " + id + " successfully updated.",
        })
    })
    router.Run(":8080")
}
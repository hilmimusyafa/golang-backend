package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Defining a route with the dynamic parameter ":id"
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id") // Retrieves the "id" parameter value from the URL
        c.JSON(http.StatusOK, gin.H{
            "message": "Detail user dengan ID: " + id,
        })
    })

    r.Run(":8080")
}
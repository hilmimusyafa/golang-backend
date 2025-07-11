package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Define a route with the dynamic parameter ":username"
    r.GET("/profile/:username", func(c *gin.Context) {
        username := c.Param("username")
        c.JSON(http.StatusOK, gin.H{
            "message": "Profile with username : " + username,
        })
    })

    r.Run(":8080")
}
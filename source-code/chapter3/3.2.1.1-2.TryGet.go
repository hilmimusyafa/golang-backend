package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/welcome", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to Server Gin!",
        })
    })

    router.GET("/checkserver", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Server is running well!",
        })
    })
    
    router.Run(":8080")
}
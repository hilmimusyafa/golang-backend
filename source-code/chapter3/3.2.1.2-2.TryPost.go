package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.POST("/create-user", func(c *gin.Context) {
        c.JSON(http.StatusCreated, gin.H{
            "message": "User created successfully.",
        })
    })

    router.POST("/create-product", func(c *gin.Context) {
        c.JSON(http.StatusCreated, gin.H{
            "message": "Product created successfully.",
        })
    })
    
    router.Run(":8080")
}
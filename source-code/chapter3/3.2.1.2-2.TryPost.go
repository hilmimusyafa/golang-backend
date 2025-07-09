package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.POST("/create-user", func(c *gin.Context) {
        c.JSON(http.StatusCreated, gin.H{
            "message": "User berhasil dibuat.",
        })
    })

    router.POST("/create-product", func(c *gin.Context) {
        c.JSON(http.StatusCreated, gin.H{
            "message": "Produk berhasil dibuat.",
        })
    })
    
    router.Run(":8080")
}
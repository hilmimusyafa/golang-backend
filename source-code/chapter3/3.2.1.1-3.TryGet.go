package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Mendefinisikan route dengan parameter dinamis :id
    router.GET("/users/:id", func(c *gin.Context) {
        // Mengambil nilai dari parameter 'id' dari URL
        id := c.Param("id")
        c.JSON(http.StatusOK, gin.H{
            "message": "Mengambil data user dengan ID: " + id,
        })
    })
    router.Run(":8080")
}
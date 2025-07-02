package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    // Inisialisasi Gin dengan middleware default (Logger dan Recovery)
    r := gin.Default()

    // Definisikan route untuk path root ("/") dengan metode GET
    r.GET("/", func(c *gin.Context) {
        // Kirim respons JSON dengan status 200 OK
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello From Gin!",
        })
    })

    // Jalankan server di port 8080
    r.Run(":8080")
}
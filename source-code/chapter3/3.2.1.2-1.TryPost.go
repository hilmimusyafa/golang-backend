package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    // Mendefinisikan route untuk metode POST di path "/create-user"
    router.POST("/create-user", func(c *gin.Context) {
        // ... Logika untuk membuat user baru akan ada di sini ...
        // Untuk saat ini, kita hanya mengirim respons konfirmasi
        c.JSON(http.StatusCreated, gin.H{
            "message": "User berhasil dibuat.",
        })
    })
    router.Run(":8080")
}
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Menyajikan seluruh isi folder "3.2.4-StaticAssets" pada prefix "/static"
    r.Static("/static", "./3.2.4-StaticAssets")

    // Menyajikan satu file statis, misal favicon
    r.StaticFile("/favicon.ico", "./3.2.4-StaticAssets/favicon.ico")

    // Endpoint biasa
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Home page"})
    })

    r.Run(":8080")
}
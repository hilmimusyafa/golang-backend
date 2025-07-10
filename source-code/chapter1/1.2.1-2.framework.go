package main

import (
    "log"
    
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default() 

    router.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello World from Gin!",
        })
    })

    log.Println("Gin Server run on :8080")
    router.Run(":8080")
}
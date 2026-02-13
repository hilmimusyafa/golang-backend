package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Router Gin Initialization
	// gin.Default() creates a router instance with default middleware:
	// - Logger (logs requests to the console)
	// - Recovery (prevents the server from crashing/panicking on fatal errors)
	r := gin.Default()

	// 2. Route Definition
	// When a user accesses GET on the root url "/", run this function
	r.GET("/", func(c *gin.Context) {
		// c.JSON will automatically convert a map/struct to JSON format
		// and set the Content-Type to application/json
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello! This is my first Gin Application",
			"status":  "success",
		})
	})

	// Additional route with parameter
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 3. Running the Server
	// By default it will run on port 8080
	// Can be changed with r.Run(":3000")
	r.Run() 
}
package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS configuration for production
	r.Use(cors.New(cors.Config{
		// Domain yang diizinkan mengakses API
		AllowOrigins: []string{
			"https://myapp.com",
			"https://www.myapp.com",
			"https://admin.myapp.com",
		},

		// Allowed HTTP methods
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},

		// Permitted headers from the client
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Request-ID",
		},

		// Headers that can be accessed by the client from the response
		ExposeHeaders: []string{
			"Content-Length",
			"X-Request-ID",
		},

		// Allow credentials (cookies, authorization headers)
		AllowCredentials: true,

		// How long can the browser cache preflight requests (OPTIONS)
		MaxAge: 12 * time.Hour,
	}))

	// Routes
	r.GET("/api/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Users retrieved successfully",
			"data":    []string{"User 1", "User 2"},
		})
	})

	r.POST("/api/users", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
		})
	})

	r.Run(":8080")
}
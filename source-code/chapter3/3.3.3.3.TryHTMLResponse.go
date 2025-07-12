package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load HTML templates from the "templates" directory
	// Gin will parse all files ending with .html in the "templates" folder
	r.LoadHTMLGlob("3.3.2-HTML/*")

	// Endpoint to render an HTML template
	r.GET("/html-page", func(c *gin.Context) {
		// Render "index.html" template with data
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Welcome to Gin HTML",
			"message": "Hello from Gin-Gonic!",
		})
	})

	// Serve a simple static HTML file (if needed, though Static is better for folders)
	r.GET("/static-html", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<h1>This is a Static HTML Page</h1><p>Served directly as raw data.</p>`))
	})

	r.Run(":8080")
}
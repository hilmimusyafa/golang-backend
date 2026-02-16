package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Response is standard wrapper for all output API
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty: sembunyikan field jika nil
}

// ErrorResponse is special structure for detailed error responses
type ErrorResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code,omitempty"`
	Details   interface{} `json:"details,omitempty"`
}

// Helper function for respon sukses
func SendSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Helper function for simple error responds
func SendError(c *gin.Context, status int, msg string) {
	c.JSON(status, Response{
		Success: false,
		Message: msg,
		Data:    nil, // Data kosong saat error
	})
}

// Helper function for respon detailed error
func SendDetailedError(c *gin.Context, status int, msg string, errorCode string, details interface{}) {
	c.JSON(status, ErrorResponse{
		Success:   false,
		Message:   msg,
		ErrorCode: errorCode,
		Details:   details,
	})
}

// Data user example user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	r := gin.Default()

	// Success endpoint with data
	r.GET("/users/:id", func(c *gin.Context) {
		// Data user simple simulation
		user := User{
			ID:    1,
			Name:  "Hilmi",
			Email: "hilmi@mail.com",
		}

		// SendSuccess helper used
		SendSuccess(c, "User retrieved successfully", user)
	})

	// Simple error endpoint
	r.GET("/users/:id/posts", func(c *gin.Context) {
		// Resources not found simulation
		SendError(c, http.StatusNotFound, "Posts not found for this user")
	})

	// Endpoint error with detail
	r.POST("/register", func(c *gin.Context) {
		// Validation error simulation
		validationErrors := map[string]string{
			"email":    "Email format invalid",
			"password": "Password must be at least 8 characters",
		}

		SendDetailedError(
			c,
			http.StatusBadRequest,
			"Validation failed",
			"VALIDATION_ERROR",
			validationErrors,
		)
	})

	// Success endpoint without data
	r.DELETE("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, Response{
			Success: true,
			Message: "User deleted successfully",
			Data:    nil,
		})
	})

	r.Run(":8080")
}
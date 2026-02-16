package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set limit memorif for upload form (default 32 MiB)
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	// 1. Single File Upload
	r.POST("/upload/single", func(c *gin.Context) {
		// Retrieving a file from a form field named "file"
		file, err := c.FormFile("avatar")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File avatar wajib diupload"})
			return
		}

		// Get the original file name
		filename := filepath.Base(file.Filename)
		
		// Specify the save location (make sure the 'uploads' folder already exists or use os.Mkdir)
		// Here we save it in the project root for simplification
		dst := "./" + filename

		// Save file to server
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "File berhasil diupload",
			"filename": filename,
			"size":     file.Size,
		})
	})

	// 2. Multiple File Upload
	r.POST("/upload/multiple", func(c *gin.Context) {
		// Parse form data first for multipart
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Gagal membaca form"})
			return
		}

		// Get file list from "photos" field
		files := form.File["photos"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			
			// Save each file
			if err := c.SaveUploadedFile(file, "./"+filename); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Gagal upload %s", filename)})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d file berhasil diupload", len(files)),
		})
	})

	r.Run(":8080")
}
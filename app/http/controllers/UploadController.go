package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadController struct{}

func NewUploadController() *UploadController {
	return &UploadController{}
}

// UploadNewsImage handles news image uploads
func (ctrl UploadController) UploadNewsImage(c *gin.Context) {
	// Get file from form
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "No file uploaded", "data": nil})
		return
	}

	// Validate file type
	allowedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedTypes[ext] {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Invalid file type. Only jpg, jpeg, png, gif, webp allowed", "data": nil})
		return
	}

	// Validate file size (max 5MB)
	if file.Size > 5*1024*1024 {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "File too large. Max 5MB", "data": nil})
		return
	}

	// Create upload directory if not exists
	uploadDir := "./upload/news"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Failed to create upload directory", "data": nil})
		return
	}

	// Generate unique filename with timestamp and UUID
	timestamp := time.Now().Format("20060102_150405")
	uniqueID := uuid.New().String()[:8]
	newFilename := fmt.Sprintf("%s_%s%s", timestamp, uniqueID, ext)
	fullPath := filepath.Join(uploadDir, newFilename)

	// Save file
	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Failed to save file: " + err.Error(), "data": nil})
		return
	}

	// Return full path
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "File uploaded successfully",
		"data": gin.H{
			"filename": newFilename,
			"path":     fullPath,
			"url":      "/upload/news/" + newFilename,
		},
	})
}

// DeleteNewsImage handles news image deletion
func (ctrl UploadController) DeleteNewsImage(c *gin.Context) {
	var input struct {
		Path string `json:"path" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
		return
	}

	// Validate path is in allowed directory
	if !strings.HasPrefix(input.Path, "./upload/news/") {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Invalid file path", "data": nil})
		return
	}

	// Delete file
	if err := os.Remove(input.Path); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Failed to delete file: " + err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "File deleted successfully", "data": nil})
}

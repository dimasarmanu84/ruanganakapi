package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type MediaController struct {
	db *gorm.DB
}

func NewMediaController() *MediaController {
	return &MediaController{
		db: app.DB,
	}
}

func (ctrl MediaController) Edit(c *gin.Context) {
	id := c.Param("id")

	var media models.Media

	if err := ctrl.db.Select("*").Table("file_trx_media").Where("file_id::text = ?", id).Scan(&media).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if strings.Contains(media.FileType, "image") {
		file, err := os.Open(media.FileFullPath)
		if err != nil {
			c.Error(err)
			return
		}
		defer file.Close()

		// Set the content type and headers
		c.Header("Content-Type", "image/jpeg")
		c.Header("Cache-Control", "no-cache")

		// Stream the image
		io.Copy(c.Writer, file)
		return
	} else {

		file, err := os.Open(media.FileFullPath)
		if err != nil {
			c.Error(err)
			return
		}

		defer file.Close()

		c.Header("Content-Type", "video/mp4")
		buffer := make([]byte, 64*1024) // 64KB buffer size
		io.CopyBuffer(c.Writer, file, buffer)
		return
	}
}

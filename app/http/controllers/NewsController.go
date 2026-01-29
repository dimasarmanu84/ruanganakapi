package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	pagination "github.com/Hironaga06/gorm-pagination"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type NewsController struct {
	db *gorm.DB
}

func NewNewsController() *NewsController {
	return &NewsController{
		db: app.DB,
	}
}

// DataTable handles paginated listing of news
func (ctrl NewsController) DataTable(c *gin.Context) {
	var news []models.AppTrxNews

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Invalid offset", "data": nil})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Invalid limit", "data": nil})
		return
	}

	search := c.DefaultQuery("search", "")

	query := ctrl.db.Select("*").Table("app_trx_news")

	if search != "" {
		query = query.Where("LOWER(news_title) LIKE ? OR LOWER(news_desc) LIKE ?",
			"%"+search+"%", "%"+search+"%")
	}

	query.Order("news_publish DESC")
	query.Find(&news)

	if query.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": query.Error.Error(), "data": nil})
		return
	}

	p := pagination.New(query, offset, limit, []string{"news_publish desc"}, &news, true)
	result, err := p.Paging()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAll returns all news without pagination
func (ctrl NewsController) GetAll(c *gin.Context) {
	var news []models.AppTrxNews

	limit := c.DefaultQuery("limit", "100")

	query := ctrl.db.Select("*").Table("app_trx_news").
		Order("news_publish DESC").
		Limit(100)

	if limit != "" {
		if limitInt, err := strconv.Atoi(limit); err == nil {
			query = query.Limit(limitInt)
		}
	}

	if err := query.Find(&news).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": news})
}

// GetByID returns a single news item by ID
func (ctrl NewsController) GetByID(c *gin.Context) {
	id := c.Param("id")
	var news models.AppTrxNews

	if err := ctrl.db.Where("news_id = ?", id).First(&news).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Record not found", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": news})
}

// Create inserts a new news item
func (ctrl NewsController) Create(c *gin.Context) {
	var input models.AppTrxNews

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
		return
	}

	// Generate UUID if not provided
	if input.NewsID == "" {
		input.NewsID = uuid.New().String()
	}

	if err := ctrl.db.Create(&input).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "News created successfully", "data": input})
}

// Update modifies an existing news item
func (ctrl NewsController) Update(c *gin.Context) {
	id := c.Param("id")
	var news models.AppTrxNews

	// Find existing record
	if err := ctrl.db.Where("news_id = ?", id).First(&news).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Record not found", "data": nil})
		return
	}

	// Use raw map to get all fields
	var rawInput map[string]interface{}
	if err := c.ShouldBindJSON(&rawInput); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
		return
	}

	// Debug
	fmt.Println("=== Update News Debug ===")
	fmt.Println("News ID:", id)
	for k, v := range rawInput {
		fmt.Printf("%s: %v (type: %T)\n", k, v, v)
	}

	// Build update map with proper column names
	updates := make(map[string]interface{})

	if val, ok := rawInput["news_title"]; ok {
		updates["news_title"] = val
	}
	if val, ok := rawInput["news_desc"]; ok {
		updates["news_desc"] = val
	}
	if val, ok := rawInput["news_link"]; ok {
		updates["news_link"] = val
	}
	if val, ok := rawInput["news_photo"]; ok {
		updates["news_photo"] = val
	}
	if val, ok := rawInput["news_publish"]; ok {
		updates["news_publish"] = val
	}
	if val, ok := rawInput["news_unpublish"]; ok {
		updates["news_unpublish"] = val
	}
	if val, ok := rawInput["school_id"]; ok {
		updates["school_id"] = val
	}

	// Handle is_active separately with raw SQL to ensure it updates correctly
	var isActiveUpdated bool
	var isActiveValue interface{}
	if val, ok := rawInput["is_active"]; ok {
		isActiveValue = val
		isActiveUpdated = true
		fmt.Printf("Will update is_active to: %v\n", val)
	}

	fmt.Println("========================")

	// Update other fields first (excluding is_active)
	if len(updates) > 0 {
		if err := ctrl.db.Model(&news).UpdateColumns(updates).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
			return
		}
	}

	// Update is_active separately using Exec to force update
	if isActiveUpdated {
		sql := "UPDATE app_trx_news SET is_active = ? WHERE news_id = ?"
		if err := ctrl.db.Exec(sql, isActiveValue, id).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": true, "message": "Failed to update is_active: " + err.Error(), "data": nil})
			return
		}
		fmt.Printf("Executed raw SQL to update is_active to: %v\n", isActiveValue)
	}

	// Fetch the updated record to confirm
	if err := ctrl.db.Where("news_id = ?", id).First(&news).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Failed to fetch updated record", "data": nil})
		return
	}

	if news.IsActive != nil {
		fmt.Printf("After update - IsActive: %v (actual value)\n", *news.IsActive)
	} else {
		fmt.Println("After update - IsActive: nil")
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "News updated successfully", "data": news})
}

// Delete removes a news item
func (ctrl NewsController) Delete(c *gin.Context) {
	id := c.Param("id")
	var news models.AppTrxNews

	if err := ctrl.db.Where("news_id = ?", id).First(&news).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Record not found", "data": nil})
		return
	}

	if err := ctrl.db.Delete(&news).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "News deleted successfully", "data": nil})
}

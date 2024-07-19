package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type InvoiceController struct {
	db *gorm.DB
}

func NewInvoiceController() *InvoiceController {
	return &InvoiceController{
		db: app.DB,
	}
}

func (ctrl InvoiceController) Generate(c *gin.Context) {

	var child []models.ViewChild
	var branch models.ViewBranch
	var school models.ViewSchool

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	postdata := JsonToArray(jsonData)
	if err := ctrl.db.Table("peo_vw_child").Where("branch_id::text = ?", postdata["branch_id"]).Scan(&child).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "Record  child not found!": err.Error(), "data": nil})
		return
	}

	if err := ctrl.db.Table("sch_vw_branch").Where("branch_id::text = ?", postdata["branch_id"]).First(&branch).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "Record branch not found!": err.Error(), "data": nil})
		return
	}

	if err := ctrl.db.Table("sch_vw_school").Where("school_id::text = ?", branch.SchoolId).First(&school).Error; err != nil {

		c.JSON(http.StatusOK, gin.H{"error": true, "Record school not found!": err.Error(), "data": nil})
		return
	}

	Result := map[string]interface{}{
		"child":  &child,
		"school": &school,
		"branch": &branch,
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &Result})

}

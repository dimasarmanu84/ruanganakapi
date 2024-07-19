package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type MilestoneController struct {
	db *gorm.DB
}

func NewMilestoneController() *MilestoneController {
	return &MilestoneController{
		db: app.DB,
	}
}

func (ctrl MilestoneController) Edit(c *gin.Context) {
	var milestone []models.ViewMilestoneGroup
	var child models.ViewChild
	var branch models.ViewBranch
	var school models.ViewSchool

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	postdata := JsonToArray(jsonData)

	if err := ctrl.db.Table("peo_vw_child").Where("child_id::text = ?", postdata["child_id"]).First(&child).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "Record  child not found!": err.Error(), "data": nil})
		return
	}

	if err := ctrl.db.Table("sch_vw_branch").Where("branch_id::text = ?", child.BranchId).First(&branch).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "Record branch not found!": err.Error(), "data": nil})
		return
	}

	if err := ctrl.db.Table("sch_vw_school").Where("school_id::text = ?", branch.SchoolId).First(&school).Error; err != nil {

		c.JSON(http.StatusOK, gin.H{"error": true, "Record school not found!": err.Error(), "data": nil})
		return
	}

	// var query *gorm.DB

	if err := ctrl.db.Select("*").Table("mile_vw_milestone_group").Where("child_id::text = ?", postdata["child_id"]).Where("mile_vw_milestone_group.report_date BETWEEN ? AND ?", postdata["start_date"], postdata["end_date"]).Scan(&milestone).Error; err != nil {

		c.JSON(http.StatusOK, gin.H{"error": true, "Record milestone not found!": err.Error(), "data": nil})
		return
	}

	Result := map[string]interface{}{
		"child":     &child,
		"school":    &school,
		"branch":    &branch,
		"milestone": &milestone,
		"postdata":  &postdata,
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &Result})

}

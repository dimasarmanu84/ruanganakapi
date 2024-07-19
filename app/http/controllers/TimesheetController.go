package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type TimesheetController struct {
	db *gorm.DB
}

func NewTimesheetController() *TimesheetController {
	return &TimesheetController{
		db: app.DB,
	}
}

func (ctrl TimesheetController) Edit(c *gin.Context) {
	var timesheet []models.ViewChildTimesheetOvertime
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record child not found!"})
		return
	}

	if err := ctrl.db.Table("sch_vw_branch").Where("branch_id::text = ?", child.BranchId).First(&branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record branch not found!"})
		return
	}

	if err := ctrl.db.Table("sch_vw_school").Where("school_id::text = ?", branch.SchoolId).First(&school).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record school not found!"})
		return
	}

	// var query *gorm.DB

	if err := ctrl.db.Select("*, to_char(clock_in, 'HH:mi:ss') as clock_in,to_char(clock_out, 'HH:mi:ss') as clock_out").Table("v_timesheet_child_overtime").Where("child_id::text = ?", postdata["child_id"]).Where("v_timesheet_child_overtime.date BETWEEN ? AND ?", postdata["start_date"], postdata["end_date"]).Scan(&timesheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record milestone not found!"})
		return
	}

	Result := map[string]interface{}{
		"child":     &child,
		"school":    &school,
		"branch":    &branch,
		"timesheet": &timesheet,
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &Result})
}

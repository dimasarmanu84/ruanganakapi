package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type TimesheetEducatorController struct {
	db *gorm.DB
}

func NewTimesheetEducatorController() *TimesheetEducatorController {
	return &TimesheetEducatorController{
		db: app.DB,
	}
}

func (ctrl TimesheetEducatorController) Edit(c *gin.Context) {
	response := new(models.Response)
	var timesheeteducator []models.ViewEducatorTimesheet
	var educator models.ViewEducator

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	postdata := JsonToArray(jsonData)

	// var query *gorm.DB

	if err := ctrl.db.Table("peo_vw_educator").Where("user_id::text = ?", postdata["user_id"]).First(&educator).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := ctrl.db.Select("*, to_char(clock_in, 'HH:mi:ss') as clock_in,to_char(clock_out, 'HH:mi:ss') as clock_out").Table("time_vw_educator_timesheet").Where("user_id::text = ?", postdata["user_id"]).Where("time_vw_educator_timesheet.date BETWEEN ? AND ?", postdata["start_date"], postdata["end_date"]).Scan(&timesheeteducator).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	Result := map[string]interface{}{
		"educator":  &educator,
		"timesheet": &timesheeteducator,
	}

	response.Status = http.StatusOK // if the post is created successfully, return a success response
	response.Message = "Success"
	response.Data = &Result
	c.JSON(200, response)
}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	pagination "github.com/Hironaga06/gorm-pagination"
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type SchoolController struct {
	db *gorm.DB
}

func NewSchoolController() *SchoolController {
	return &SchoolController{
		db: app.DB,
	}
}

func (ctrl SchoolController) DataTable(c *gin.Context) {

	var school []models.SchMstSchool

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil {
		c.JSON(200, nil)
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		c.JSON(200, nil)
		return
	}

	search := c.DefaultQuery("search", "")

	var query *gorm.DB
	query = ctrl.db.Select("*").Table("sch_mst_school")

	if search != "" {
		query = query.Where("LOWER(sch_mst_school.school_name) LIKE  ?", "%"+search+"%")
	}

	// if datatable.Search2 != "" {
	// 	query = query.Or("tmp_mst_role.description ILIKE  ?", "%"+datatable.Search2+"%")
	// }

	query.Find(&school)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	p := pagination.New(query, offset, limit, []string{"school_id asc"}, &school, true)
	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

func (ctrl SchoolController) Update(c *gin.Context) {
	id := c.Param("id")
	var school models.SchMstSchool

	if err := ctrl.db.Where("school_id = ?", id).First(&school).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var input models.SchMstSchool
	// Parse the updated fields from the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctrl.db.Model(&school).Where("school_id", id).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &school})
}

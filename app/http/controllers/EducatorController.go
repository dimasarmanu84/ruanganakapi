package controllers

import (
	"fmt"

	pagination "github.com/Hironaga06/gorm-pagination"
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type EducatorController struct {
	db *gorm.DB
}

func NewEducatorController() *EducatorController {
	return &EducatorController{
		db: app.DB,
	}
}

func (ctrl EducatorController) DataTable(c *gin.Context) {

	var educator []models.ViewEducator

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	datatable := JsonToArray(jsonData)

	var query *gorm.DB

	query = ctrl.db.Select("*").Table("peo_vw_educator")

	if datatable["search"] != nil {
		query = query.Where("LOWER(full_name) ILIKE  ?", "%"+datatable["search"].(string)+"%")
	}

	if datatable["childname"] != nil {
		query = query.Where("LOWER(full_name) ILIKE  ?", "%"+datatable["childname"].(string)+"%")
	}

	if datatable["branchid"] != nil {
		query = query.Where("branch_id::TEXT ILIKE ?", "%"+datatable["branchid"].(string)+"%")
	}

	query.Find(&educator)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	p := pagination.New(query, int(datatable["offset"].(float64)), int(datatable["limit"].(float64)), []string{"educator_id asc"}, &educator, true)
	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

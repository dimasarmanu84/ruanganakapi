package controllers

import (
	"fmt"

	pagination "github.com/Hironaga06/gorm-pagination"
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type ParentsController struct {
	db *gorm.DB
}

func NewParentsController() *ParentsController {
	return &ParentsController{
		db: app.DB,
	}
}

func (ctrl ParentsController) DataTable(c *gin.Context) {

	var parents []models.ViewParents

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	datatable := JsonToArray(jsonData)

	var query *gorm.DB

	query = ctrl.db.Select("*").Table("peo_vw_parent")

	if datatable["search"] != nil {
		query = query.Where("LOWER(full_name) ILIKE  ?", "%"+datatable["search"].(string)+"%")
	}

	if datatable["childname"] != nil {
		query = query.Where("LOWER(full_name) ILIKE  ?", "%"+datatable["full_name"].(string)+"%")
	}

	if datatable["branch_id"] != nil {
		query = query.Where("branch_id::TEXT ILIKE ?", "%"+datatable["branch_id"].(string)+"%")
	}

	query.Find(&parents)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	p := pagination.New(query, int(datatable["offset"].(float64)), int(datatable["limit"].(float64)), []string{"user_id asc"}, &parents, true)
	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

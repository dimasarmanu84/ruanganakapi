package controllers

import (
	"fmt"
	"net/http"

	pagination "github.com/Hironaga06/gorm-pagination"
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type ChildsController struct {
	db *gorm.DB
}

func NewChildsController() *ChildsController {
	return &ChildsController{
		db: app.DB,
	}
}

func (ctrl ChildsController) DataTable(c *gin.Context) {

	var childs []models.ViewChild

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	datatable := JsonToArray(jsonData)

	var query *gorm.DB

	query = ctrl.db.Select("*,to_char(child_dob, 'DD-MM-YYYY') as child_dob").Table("peo_vw_child")

	if datatable["search"] != nil {
		query = query.Where("LOWER(child_full_name) ILIKE  ?", "%"+datatable["search"].(string)+"%")
	}

	if datatable["childname"] != nil {
		query = query.Where("LOWER(child_full_name) ILIKE  ?", "%"+datatable["childname"].(string)+"%")
	}

	if datatable["branch_id"] != nil {
		query = query.Where("branch_id::TEXT ILIKE ?", "%"+datatable["branch_id"].(string)+"%")
	}

	query.Find(&childs)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	p := pagination.New(query, int(datatable["offset"].(float64)), int(datatable["limit"].(float64)), []string{"child_id asc"}, &childs, true)
	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

func (ctrl ChildsController) GetChildByUserid(c *gin.Context) {
	response := new(models.Response)
	var child []models.ViewChild

	if err := ctrl.db.Select("*,to_char(child_dob, 'DD-MM-YYYY') as child_dob").Table("peo_vw_child").Where("user_id::text = ?", c.Param("id")).Find(&child).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	response.Status = http.StatusOK // if the post is created successfully, return a success response
	response.Message = "Success"
	response.Data = &child
	c.JSON(200, response)
}

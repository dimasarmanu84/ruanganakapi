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

type BranchController struct {
	db *gorm.DB
}

func NewBranchController() *BranchController {
	return &BranchController{
		db: app.DB,
	}
}

func (ctrl BranchController) DataTable(c *gin.Context) {

	var branch []models.SchMstBranch

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	datatable := JsonToArray(jsonData)

	var query *gorm.DB

	query = ctrl.db.Select("*").Table("sch_mst_branch")

	if datatable["search"] != nil {
		query = query.Where("LOWER(sch_mst_branch.branch_name) LIKE  ?", "%"+datatable["search"].(string)+"%")
	}

	query.Find(&branch)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}
	p := pagination.New(query, int(datatable["offset"].(float64)), int(datatable["limit"].(float64)), []string{"branch_id asc"}, &branch, true)

	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

func (ctrl BranchController) Dropdown(c *gin.Context) {

	var results []map[string]interface{}

	var query *gorm.DB
	query = ctrl.db.Select("sch_mst_branch.branch_id AS value, sch_mst_branch.branch_name AS name").Table("sch_mst_branch").Find(&results)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	c.JSON(200, results)
}

func (ctrl BranchController) Update(c *gin.Context) {
	id := c.Param("id")
	var branch models.SchMstBranch

	if err := ctrl.db.Where("school_id = ?", id).First(&branch).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var input models.SchMstBranch
	// Parse the updated fields from the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctrl.db.Model(&branch).Where("school_id", id).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &branch})
}

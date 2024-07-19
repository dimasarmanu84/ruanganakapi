package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	pagination "github.com/Hironaga06/gorm-pagination"
	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type RolesController struct {
	db *gorm.DB
}

func NewRolesController() *RolesController {
	return &RolesController{
		db: app.DB,
	}
}

func (ctrl RolesController) DataTable(c *gin.Context) {

	var role []models.Roles
	var datatable models.DataTables

	// This reads c.Request.Body and stores the result into the context.
	if err := c.ShouldBindJSON(&datatable); err != nil {
		c.JSON(200, nil)
		return
	}

	var query *gorm.DB
	query = ctrl.db.Select("*").Table("dss_main.tmp_mst_role").Where("is_active =?", "Y")

	if datatable.Search != "" {
		query = query.Where("tmp_mst_role.name ILIKE  ?", "%"+datatable.Search+"%")
	}

	if datatable.Search2 != "" {
		query = query.Or("tmp_mst_role.description ILIKE  ?", "%"+datatable.Search2+"%")
	}

	query.Find(&role)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	p := pagination.New(query, datatable.Offset, datatable.Limit, []string{"role_id asc"}, &role, true)
	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

func (ctrl RolesController) Edit(c *gin.Context) {
	response := new(models.Response)
	var role models.Roles

	if err := ctrl.db.Table("dss_main.tmp_mst_role").Where("role_id = ?", c.Param("id")).First(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	response.Status = http.StatusOK // if the post is created successfully, return a success response
	response.Message = "Success"
	response.Data = &role
	c.JSON(200, response)
}

func (ctrl RolesController) Create(c *gin.Context) {

	response := new(models.Response) // initialize the response variable
	role := new(models.Roles)

	if err := c.ShouldBindJSON(&role); err == nil {
		ctrl.db.Create(&role)
		response.Status = http.StatusOK // if the post is created successfully, return a success response
		response.Message = "Success"
		response.Data = &role
		c.JSON(200, response)
	} else {
		response.Status = http.StatusNotFound // if the post is created successfully, return a success response
		response.Message = err.Error()
		response.Data = nil
		c.JSON(200, response)
	}
}

func (ctrl RolesController) Update(c *gin.Context) {
	var role models.Roles
	response := new(models.Response)

	id := c.Param("id")

	if err := ctrl.db.Table("dss_main.tmp_mst_role").Where("role_id = ?", id).First(&role).Error; err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Error"
		response.Data = nil
		c.JSON(200, response)
		return
	}

	var input models.Roles
	// Parse the updated fields from the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	}
	ctrl.db.Table("dss_main.tmp_mst_role").Where("role_id", id).Updates(&input)

	// Return a success response
	response.Status = http.StatusOK // if the post is created successfully, return a success response
	response.Message = "Success"
	response.Data = &role
	c.JSON(200, response)

}

func (ctrl RolesController) Test(c *gin.Context) {
	mapData := map[string]interface{}{
		"Name":    "noknow",
		"Age":     2,
		"Admin":   true,
		"Hobbies": []string{"IT", "Travel"},
		"Address": map[string]interface{}{
			"PostalCode": 1111,
			"Country":    "Japan",
		},
		"Null": nil,
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(mapData)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, jsonStr)

}

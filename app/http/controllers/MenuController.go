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

type MenusController struct {
	db *gorm.DB
}

// This is a comment for the new function
func NewMenuController() *MenusController {
	return &MenusController{
		db: app.DB,
	}
}

func (ctrl MenusController) GetParentNotId(c *gin.Context) {
	response := new(models.Response)

	id, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	query := `SELECT dss_main.tmp_mst_menu.menu_id AS value, dss_main.tmp_mst_menu.name as name
			FROM  dss_main.tmp_mst_menu
			WHERE dss_main.tmp_mst_menu.menu_id <> ? ORDER BY dss_main.tmp_mst_menu.name ASC`

	var results []map[string]interface{}
	err := NewMenuController().db.Raw(query, id).Scan(&results).Error

	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Error"
		response.Data = nil
		c.JSON(200, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = &results
	c.JSON(200, response)
}

func (ctrl MenusController) GetMenuByUserid(c *gin.Context) {
	var menus []map[string]interface{}

	response := new(models.Response)

	id, exists := c.Params.Get("id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}
	menus, err := GetMenus(id, "0")

	for _, menu := range menus {
		var o []map[string]interface{}
		str := fmt.Sprint(menu["menu_id"])

		o, err := GetMenus(id, str)
		if err != nil {
			response.Status = http.StatusBadRequest
			response.Message = "Error"
			response.Data = nil
			c.JSON(200, response)
			return
		}
		if len(o) > 0 {
			menu["sub_menu"] = o
		}

		for _, p := range o {
			var d []map[string]interface{}
			str2 := fmt.Sprint(p["menu_id"])
			d, e := GetMenus(id, str2)

			if e != nil {
				response.Status = http.StatusBadRequest
				response.Message = "Error"
				response.Data = nil
				c.JSON(200, response)
				return
			}

			if len(d) > 0 {
				p["sub_menu"] = d
			}
		}

		//fmt.Printf("%v", o)
	}

	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Error"
		response.Data = nil
		c.JSON(200, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = &menus
	c.JSON(200, response)
}

func GetMenus(id string, parent string) ([]map[string]interface{}, error) {

	query := `SELECT 
				dss_main.tmp_role_has_menu.menu_id,
				dss_main.tmp_mst_menu.name AS menu_name,
				COALESCE(dss_main.tmp_mst_menu.class, '') as class,
				dss_main.tmp_mst_menu.link,
				dss_main.tmp_mst_menu.parent,
				COALESCE(dss_main.tmp_mst_menu.icon, '') as icon,
				dss_main.tmp_mst_menu.counter,
				dss_main.tmp_role_has_menu.role_id,
				dss_main.tmp_mst_menu.display_on_tree 
			FROM 
				dss_main.tmp_role_has_menu 
			JOIN 
				dss_main.tmp_mst_menu 
			ON 
				dss_main. tmp_mst_menu.menu_id=dss_main.tmp_role_has_menu.menu_id 
			join 
				dss_main.tmp_user on dss_main.tmp_role_has_menu.role_id  = dss_main.tmp_user.role_id 
			where (dss_main.tmp_user.user_id = ?) AND (dss_main.tmp_mst_menu.is_active = 'Y' AND dss_main.tmp_mst_menu.parent = ?) 
			GROUP BY 
				dss_main.tmp_role_has_menu.menu_id,
				dss_main.tmp_mst_menu.name,
				dss_main.tmp_mst_menu.class,
				dss_main.tmp_mst_menu.link,
				dss_main.tmp_mst_menu.parent,
				dss_main.tmp_mst_menu.icon,
				dss_main.tmp_mst_menu.counter,
				dss_main.tmp_role_has_menu.role_id, 
				dss_main.tmp_mst_menu.display_on_tree
			ORDER BY 
				dss_main.tmp_mst_menu.counter ASC`

	var results []map[string]interface{}
	err := NewMenuController().db.Raw(query, id, parent).Scan(&results).Error
	//fmt.Printf("%v", menus)

	return results, err
}

func (ctrl MenusController) DataTable(c *gin.Context) {

	var menu []models.Menu

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
	if search == "" {
		search = ""
	}

	query := ctrl.db.Select("tmp_mst_menu.menu_id,tmp_mst_menu.name,tmp_mst_menu.class,tmp_mst_menu.link, tmp_mst_menu.counter,tmp_mst_menu.parent,tmp_mst_menu.is_active").
		Table("dss_main.tmp_mst_menu").
		Where("tmp_mst_menu.is_active = ?", "Y").
		Where("tmp_mst_menu.name like ?", "%"+search+"%").Find(&menu).Order("tmp_mst_menu.link ASC")

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	p := pagination.New(query, offset, limit, []string{"menu_id asc"}, &menu, false)
	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

func (ctrl MenusController) Edit(c *gin.Context) {
	response := new(models.Response)
	var menu models.Menu

	if err := ctrl.db.Table("dss_main.tmp_mst_menu").Where("menu_id = ?", c.Param("id")).First(&menu).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	response.Status = http.StatusOK // if the post is created successfully, return a success response
	response.Message = "Success"
	response.Data = &menu
	c.JSON(200, response)
}

func (ctrl MenusController) Create(c *gin.Context) {

	response := new(models.Response) // initialize the response variable
	menu := new(models.Menu)

	if err := c.ShouldBindJSON(&menu); err == nil {
		ctrl.db.Create(&menu)
		response.Status = http.StatusOK // if the post is created successfully, return a success response
		response.Message = "Success"
		response.Data = &menu
		c.JSON(200, response)
	} else {
		response.Status = http.StatusNotFound // if the post is created successfully, return a success response
		response.Message = err.Error()
		response.Data = nil
		c.JSON(200, response)
	}
}

func (ctrl MenusController) Update(c *gin.Context) {
	var menu models.Menu
	response := new(models.Response)

	id := c.Param("id")

	if err := ctrl.db.Table("dss_main.tmp_mst_menu").Where("menu_id = ?", id).First(&menu).Error; err != nil {
		response.Status = http.StatusBadRequest // if the post is created successfully, return a success response
		response.Message = "Error"
		response.Data = nil
		c.JSON(200, response)
		return
	}

	var input models.Menu
	// Parse the updated fields from the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(200, gin.H{
			"error": err,
		})
		return
	}
	fmt.Printf("%v", &input)

	ctrl.db.Table("dss_main.tmp_mst_menu").Where("menu_id", id).Updates(&input)

	// Return a success response
	response.Status = http.StatusOK // if the post is created successfully, return a success response
	response.Message = "Success"
	response.Data = &menu
	c.JSON(200, response)

}

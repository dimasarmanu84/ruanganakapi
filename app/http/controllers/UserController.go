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

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

type UserController struct {
	db *gorm.DB
}

// This is a comment for the new function
func NewUserController() *UserController {
	return &UserController{
		db: app.DB,
	}
}

func (ctrl UserController) DataTable(c *gin.Context) {

	var branch []models.AppMstUser

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
	query = ctrl.db.Select("*").Table("app_mst_user")

	if search != "" {
		query = query.Where("LOWER(app_mst_user.user_phone) LIKE  ?", "%"+search+"%")
	}

	// if datatable.Search2 != "" {
	// 	query = query.Or("tmp_mst_role.description ILIKE  ?", "%"+datatable.Search2+"%")
	// }

	query.Find(&branch)

	if query.Error != nil {
		fmt.Println(query.Error)
		return
	}

	p := pagination.New(query, offset, limit, []string{"user_id asc"}, &branch, true)
	result, err := p.Paging()
	if err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, result)
}

func (ctrl UserController) Update(c *gin.Context) {
	id := c.Param("id")
	var appmstuser models.AppMstUser

	if err := ctrl.db.Where("user_id::TEXT = ?", id).First(&appmstuser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var input models.AppMstUser
	// Parse the updated fields from the request body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctrl.db.Model(&appmstuser).Where("user_id::TEXT", id).Updates(&input)
	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &appmstuser})
}

func (ctrl UserController) UpdateOTP(c *gin.Context) {
	id := c.Param("id")

	var appmstuser models.AppMstUser
	if err := ctrl.db.Where("user_phone = ?", id).First(&appmstuser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	last6Digits := id[len(id)-6:]
	codeInt, err := strconv.ParseInt(last6Digits, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid OTP"})
		return
	}

	if err = ctrl.db.Table("app_mst_user").Where("user_phone = ?", id).Updates(map[string]interface{}{"user_otp": codeInt, "user_device": nil, "user_session": nil}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": true, "message": "Success", "data": &err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &appmstuser})
}

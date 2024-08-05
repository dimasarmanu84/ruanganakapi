package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/app/models"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/bootstrap/app"
	"gorm.io/gorm"
)

type LoginController struct {
	db *gorm.DB
}

func NewLoginController() *LoginController {
	return &LoginController{
		db: app.DB,
	}
}

func (ctrl LoginController) DoLogin(c *gin.Context) {

	var useradmin models.AppMstAdminUser
	// Parse the updated fields from the request body
	if err := c.ShouldBindJSON(&useradmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	if err := ctrl.db.Table("app_mst_admin_user").Where("admin_user_name = ?", useradmin.AdminUserName).Where("admin_user_password = ?", useradmin.AdminUserPassword).First(&useradmin).Error; err != nil {
		c.JSON(200, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "message": "Success", "data": &useradmin})
}

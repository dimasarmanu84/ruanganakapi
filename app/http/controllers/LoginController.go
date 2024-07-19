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

	response := new(models.Response)

	type NamedArgument struct {
		v_user_name     string
		v_user_password string
	}

	query := "SELECT public.admin_do_login('admin','123456')"

	var results []map[string]interface{}

	err := NewMenuController().db.Raw(query).Scan(&results).Error
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Message = "Error"
		response.Data = nil
		c.JSON(200, response)
		return
	}
	c.JSON(200, results)
}

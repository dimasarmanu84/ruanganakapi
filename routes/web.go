package routes

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func SetupWebRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ginBoilerplateVersion": "v0.01",
			"goVersion":             runtime.Version(),
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":      false,
			"status_code": http.StatusForbidden,
			"message":     "Route Not Found",
		})
	})

	router.StaticFS("/images", http.Dir("./public"))

}

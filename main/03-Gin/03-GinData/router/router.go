package router

import (
	"ginRouter/service"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.POST("/login", service.CheckJson)
		userGroup.GET("/urlData/:username/:pwd", service.CheckUrlData)
		userGroup.GET("/formData", service.CheckFormData)
	}
}

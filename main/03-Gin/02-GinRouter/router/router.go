package router

import (
	"ginRouter/controller/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", initTest)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", user.LoginFunc)
		userGroup.GET("/register", user.Register)
	}
}

func initTest(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "ok!",
		"data": "this is a interface",
	})
}

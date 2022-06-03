package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", initTest)
}

func initTest(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "ok!",
		"data": "this is a interface",
	})
	
}

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required"`
	Pwd      string `form:"pwd" json:"pwd" uri:"pwd" xml:"pwd" binding:"required"`
}

func UserLogin(c *gin.Context) {
	// 对json数据进行校验
	var json Login

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "params error!",
		})
		return
	}
	fmt.Println(json)
	if json.Username == "zyd" {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "login success",
		})
		return
	}

	c.JSON(400, gin.H{
		"code": 200,
		"msg":  "login error",
	})
}

func CheckUrlData(c *gin.Context) {
	var urlParams Login
	if err := c.ShouldBindUri(&urlParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("ok！")

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "login success",
	})
}

func CheckFormData(c *gin.Context) {
	var formUrl Login
	if err := c.Bind(&formUrl); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "login error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "login success",
	})

}

func returnData(e error, c *gin.Context) {
	if e != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  "login error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "login success",
	})
}

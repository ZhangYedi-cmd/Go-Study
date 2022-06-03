# Gin数据校验



## Json数据解析和绑定

我们还是以一个登录功能作为示例
在函数中，我们使用gin框架提供的ShouldBindJSON方法校验数据，如果客户端传过来的json数据满足要求，则err为空

```go
package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义请求过来的结构体
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Pwd      string `form:"pwd" json:"pwd" binding:"required"`
}

func CheckJson(c *gin.Context) {
	// 对json数据进行校验
	var json Login
    
	// 如果请求过来的Json数据满足要求，则err不为nil
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



```



## URL数据校验


先注册路由： 
在这个模板路径中，我们接受两个参数，username和pwd
```go
userGroup.GET("/urlData/:username/:pwd", service.CheckUrlData)
```
同理，在处理函数CheckUrlData中，我们使用gin框架提供的ShouldBindUri方法校验数据

```go
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
```


## FormData 数据
同理，校验FormData格式的数据也是如此，只是换了一个方法而已。
```go
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
```
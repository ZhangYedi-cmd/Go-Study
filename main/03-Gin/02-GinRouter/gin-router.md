# Gin - 路由

## 1,路由示例


### 1. 初始化目录结构

+ config     配置选项
+ controller 控制器
+ router     路由
  + router.go 路由注册处理函数
+ main.go  程序启动文件



### 2. 在main.go中新建一个Gin服务器
```go

package main

import (
	"fmt"
	"ginRouter/router"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() // 创建一个Gin
	router.InitRouter(engine) // 注册路由
	engine.Run(":8040") // 启动服务
}

```
### 3.在router中配置注册路由函数

```go

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) { 
	// 在main.go中调用注册路由的方法
	// 路由对应的controller 
	r.GET("/", initTest)
}

func initTest(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "ok!",
		"data": "this is a interface",
	})
}

```


最后，启动main.go,这样gin的服务就跑起来了！
在浏览器中输入 127.0.0.1:8040/,就会触发视图函数initTest, 并且返回一个JSON数据

## 2,路由嵌套

在上面的代码实例中，我们在main.go中配置了engine，这就相当于是服务器对象，然后调用router.go中的initRouter函数，传入engine,
为配置路由.
在实际的开发中，我们经常需要对路由进行嵌套处理

```go

func InitRouter(r *gin.Engine) {
	r.GET("/", initTest)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", user.LoginFunc)
		userGroup.GET("/register", user.Register)
	}
}
```

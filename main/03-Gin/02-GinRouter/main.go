package main

import (
	"fmt"
	"ginRouter/router"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("system Start!") // 创建一个gin服务器
	engine := gin.Default()
	router.InitRouter(engine) // 初始化路由
	engine.Run(":8040")       // 指定端口运行服务
}

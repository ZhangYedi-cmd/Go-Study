package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Gin 洋葱模型

func MyMiddleWare1() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("ware1 before....")
		context.Next() // 执行下一个中间件
		fmt.Println("ware1 after....")
	}
}

// 中间件2

func MyMiddleWare2() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("ware2 before....")
		context.Next() // 执行下一个中间件
		fmt.Println("ware2 after....")
	}
}

func handleFunc(ctx *gin.Context) {
	fmt.Println("do func")
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func main() {
	engine := gin.Default()
	v1 := engine.Group("/v1")
	{
		v1.Use(MyMiddleWare1(), MyMiddleWare2())
		v1.GET("/test", handleFunc)
	}
	engine.Run(":8040")
}

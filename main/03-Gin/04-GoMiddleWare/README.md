## Gin中间件


###  中间件的创建
从源码的角度来解释如何创建中间件

gin.go

```go
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
   debugPrintWARNINGDefault()
   engine := New()  // 创建一个Gin实例
   engine.Use(Logger(), Recovery()) // 使用Logger和Recovery中间件
   return engine
}
```

如何自定义中间件，关键就在于这个Use方法上，我们点开源码看一看。

```go
// Use attaches a global middleware to the router. i.e. the middleware attached through Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
   engine.RouterGroup.Use(middleware...)
   engine.rebuild404Handlers()
   engine.rebuild405Handlers()
   return engine
}
```

Use方法接受了一个不定长参数，...HandlerFunc类型的中间件，HandlerFunc类型的函数默认会接受一个Gin.Context类型的参数.

```go
// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)
```

那么我们自定义中间件，只需要创建一个函数，让其返回一个类型为HandlerFunc的对象（or结果）就可以了。 

接下来我们开始创建中间件。。。。

```go
// 自定义的中间件
func MyMiddleWare1() gin.HandlerFunc {
   return func(context *gin.Context) {
      fmt.Println("ware1 before...."
      fmt.Println("ware1 after....")
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
	v1 := engine.Group("/v1") // 路由分组
	{
		v1.Use(MyMiddleWare1()) // 将中间件挂在到路由上
		v1.GET("/test", handleFunc)
	}
	engine.Run(":8040")
}
```

最后，启动服务器，看效果！

```go
ware1 before....  
do func
ware1 after....
```



### 洋葱模型

学过Node的同学一定直到一个Web框架叫[Koa](https://koa.bootcss.com/#application)， Koa中间件的实现和Gin中是一样的。 都是基于洋葱模型。![img](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/886/v2-15fd80fe2d28621f45e05e8342feb5b3_720w.jpg)

定义多个中间件，挂在到路由上

```go
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
```

运行结果

```go
ware1 before....
ware2 before....
do func
ware2 after....
ware1 after....
```


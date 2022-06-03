package main

import (
	"fmt"
	"ginRouter/router"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("system Start!")
	engine := gin.Default()
	router.InitRouter(engine)
	engine.Run(":8040")
}

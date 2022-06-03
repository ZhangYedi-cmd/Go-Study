package user

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{
		"msg":  "register success",
		"name": name,
	})
}

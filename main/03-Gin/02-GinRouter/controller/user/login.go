package user

import "github.com/gin-gonic/gin"

func LoginFunc(c *gin.Context) {
	name := c.Query("name")
	pwd := c.Query("pwd")

	if name != "" && pwd != "" {
		c.JSON(200, gin.H{
			"name": name,
			"pwd":  pwd,
		})
		return
	}

	c.JSON(400, gin.H{
		"message": "参数错误",
	})
}

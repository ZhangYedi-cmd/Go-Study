package user

import "github.com/gin-gonic/gin"

type loginResponse struct {
	string
}

func login(c *gin.Context) {
	name := c.Query("name")
	pwd := c.Query("pwd")

	if name != "" && pwd != "" {
		c.JSON(200, gin.H{
			"name": name,
			"pwd":  pwd,
		})
	}

	c.JSON(400, gin.H{
		"message": "参数错误",
	})
}

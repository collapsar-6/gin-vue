package controller

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "pong",
	})
}

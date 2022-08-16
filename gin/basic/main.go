package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/news", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "我是新闻页面")
	})
	r.Run(":8000")
}

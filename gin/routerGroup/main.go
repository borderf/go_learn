package main

import (
	"go_learn/gin/routerGroup/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "首页")
		})
		defaultRouter.GET("/news", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "新闻页面")
		})
	}
	routers.AdminRouterInit(r)
	routers.ApiRouterInit(r)
	r.Run(":8000")
}

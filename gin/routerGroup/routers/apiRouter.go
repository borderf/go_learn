package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {
	adminRouter := r.Group("/api")
	{
		adminRouter.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "api接口")
		})
		adminRouter.GET("/user", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "api接口-用户列表")
		})
		adminRouter.GET("/article", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "api接口-文章列表")
		})
	}
}

package routers

import (
	"go_learn/gin/controllers/admin"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "后台管理")
		})
		adminRouter.GET("/user", admin.UserController{}.Index)
		adminRouter.GET("/article", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "后台管理-文章列表")
		})
	}
}

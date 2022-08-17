package routers

import (
	"go_learn/gin/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/", admin.UserController{}.Index)
		adminRouter.GET("/user", admin.UserController{}.User)
		adminRouter.GET("/article", admin.UserController{}.Article)
	}
}

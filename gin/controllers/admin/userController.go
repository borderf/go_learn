package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (user UserController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "controller-首页内容")
}

func (user UserController) User(ctx *gin.Context) {
	ctx.String(http.StatusOK, "controller-用户列表")
}

func (user UserController) Article(ctx *gin.Context) {
	ctx.String(http.StatusOK, "controller-文章列表")
}

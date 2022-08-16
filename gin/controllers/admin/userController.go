package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (user UserController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "controller-用户列表")
}

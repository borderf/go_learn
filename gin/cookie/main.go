package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(ctx *gin.Context) {
		// 设置cookie
		ctx.SetCookie("username", "张三", 3600, "/", "localhost", false, true)
	})

	r.GET("/getCookie", func(ctx *gin.Context) {
		// 获取cookiego
		cookie, _ := ctx.Cookie("username")
		ctx.String(http.StatusOK, "cookie="+cookie)
	})
	r.GET("/deleteCookie", func(ctx *gin.Context) {
		// 删除cookie
		ctx.SetCookie("username", "张三", -1, "/", "localhost", false, true)
		ctx.String(http.StatusOK, "删除成功")
	})
	r.Run()
}

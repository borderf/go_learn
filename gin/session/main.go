package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 配置session中间件
	// 创建基于 cookie 的存储引擎，secret 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	// 配置session的中间件，store 是前面创建的存储引擎，我们可以替换成其他的存储引擎
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/hello", func(ctx *gin.Context) {
		// 获取session对象
		session := sessions.Default(ctx)
		// 设置session的过期时间
		session.Options(sessions.Options{
			MaxAge: 3600 * 6,
		})
		if session.Get("hello") != "world" {
			// session 设置值
			session.Set("hello", "world")
			// session 保存值，设置后必须调用
			session.Save()
		}
		ctx.JSON(http.StatusOK, gin.H{
			"hello": session.Get("hello"),
		})
	})
	r.Run()
}

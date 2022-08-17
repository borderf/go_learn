package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddleware(ctx *gin.Context) {
	fmt.Println("开始执行中间件")

	ctx.Next()

	fmt.Println("等待执行完成后中间件")
}

func callTimeMiddleware(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	duration := time.Since(start).Milliseconds()
	fmt.Printf("duration = %v\n", duration)
}
func main() {
	r := gin.Default()
	r.Use(initMiddleware)
	r.GET("/", callTimeMiddleware, func(ctx *gin.Context) {
		time.Sleep(time.Second)
		ctx.String(http.StatusOK, "gin首页")
	})
	r.Run(":8000")
}

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	r.GET("/jsonp", func(ctx *gin.Context) {
		a := &Article{
			Title:   "标题-jsonp",
			Desc:    "描述-jsonp",
			Content: "内容-jsonp",
		}
		ctx.JSONP(http.StatusOK, a)
	})
	r.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"msg":  "success",
			"code": "200",
		})
	})
	// 注意加载模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/news", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})
	r.Run(":8000")
}

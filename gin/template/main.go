package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 自定义模板函数，放在loadHtmlGlob前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "aaa",
			"msg":   "我是msg",
			"score": 89,
			"hobby": []string{"吃饭", "睡觉"},
			"date":  1660521581,
		})
	})
	r.Run(":8000")
}

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	s := t.Format("2006-01-02 15:04:05")
	return s
}

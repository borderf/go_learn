package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("admin/*")
	r.MaxMultipartMemory = 8 << 20
	adminRouter := r.Group("/admin")
	{
		adminRouter.GET("/user/add", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin/userAdd.html", gin.H{})
		})
		adminRouter.POST("/user/doUpload", func(ctx *gin.Context) {
			username := ctx.PostForm("username")
			file, err := ctx.FormFile("face")
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"msg":     "文件上传失败",
				})
			}
			// 文件名称
			dst := path.Join("./static/upload", file.Filename)
			ctx.SaveUploadedFile(file, dst)
			ctx.JSON(http.StatusOK, gin.H{
				"success":  true,
				"username": username,
				"dst":      dst,
			})
		})
	}
	r.Run()
}

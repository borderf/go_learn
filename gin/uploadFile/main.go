package main

import (
	"go_learn/gin/models"
	"net/http"
	"os"
	"path"
	"strconv"

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
			// 1、获取上传的文件
			file, err := ctx.FormFile("face")
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"success": false,
					"msg":     "文件上传失败",
				})
			}
			// 2、文件名称，是否合法 jpg png jpeg gif
			extName := path.Ext(file.Filename)
			allowExtMap := map[string]bool{
				".jpg":  true,
				".png":  true,
				".jpeg": true,
				".gif":  true,
			}

			if _, ok := allowExtMap[extName]; !ok {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "上传文件格式不合法",
				})
			}
			// 3、创建图片保存目录
			day := models.GetDay()
			dir := path.Join("./static/upload", day)
			err = os.MkdirAll(dir, 0666)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"msg": "创建文件目录失败",
				})
			}
			// 4、生成文件和文件保存的目录
			unix := models.GetUnix()
			filename := strconv.FormatInt(unix, 10) + extName
			dst := path.Join(dir, filename)
			// 5、执行上传
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

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		username := ctx.Query("username")
		age := ctx.DefaultQuery("age", "12")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
		})
	})

	r.POST("/user", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.GET("/userInfo", func(ctx *gin.Context) {
		user := &UserInfo{}
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	})
	r.Run(":8000")
}

package main

/*
 * author : xhs
 * date : 2022/5/28
 * describe : 获取url传值
 */

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("demo03/*")
	// 加载模板
	r.GET("/ping", func(c *gin.Context) {
		username := c.DefaultQuery("username", "default")
		password := c.Query("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// post演示
	r.GET("/user", func(context *gin.Context) {
		context.HTML(http.StatusOK, "user.html", gin.H{})
	})

	r.POST("/doAddUser", func(context *gin.Context) {
		usename := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"username": usename,
			"password": password,
		})
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

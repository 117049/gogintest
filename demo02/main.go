package main

/*
 * author : xhs
 * date : 2022/5/28
 * describe : 对html页面进行渲染
 */
import (
	"github.com/gin-gonic/gin"
	"gogintest/untils"
	"html/template"

	"net/http"
)

type Article struct {
	Title   string
	Content string
}

func main() {

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": untils.UnixToTime,
	}) // 自定义模板输出年月日，要放到加载模板之前

	// 加载模板
	r.LoadHTMLGlob("demo02/template/*")
	r.GET("/ping", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻",
			Content: "内容",
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
			"news":  news,
			"score": 88,
			"hobby": []string{"吃饭", "睡觉"},
			"data":  1653707918,
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

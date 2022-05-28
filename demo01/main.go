package main

/*
 * author : xhs
 * date : 2022/5/27
 * describe : 简单的web页面搭建
 */
import "github.com/gin-gonic/gin"

type Person struct {
	Name string
	Age  int
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		a := Person{
			Name: "zhangsan",
			Age:  14,
		}
		c.JSON(200, a)
	})

	// 具有回调函数，解决跨域问题
	r.GET("/jsonp", func(c *gin.Context) {
		a := Person{
			Name: "zhangsan",
			Age:  14,
		}
		c.JSONP(200, a)
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

package main

/*
 * author : xhs
 * date : 2022/5/27
 * describe : 路由分组,中间件的简易添加
 */
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func initMiddleware(context *gin.Context) {
	fmt.Println("aaaaa")
	context.Next() // 跳转执行之后再返回执行下面的语句
	fmt.Println("bbbbb")
}
func main() {

	r := gin.Default()
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", initMiddleware, func(context *gin.Context) {
			context.String(200, "首页")
		})
		defaultRouters.GET("/news", initMiddleware, func(context *gin.Context) {
			time.Sleep(time.Second * 2)
			fmt.Println("新闻展示")
			context.String(200, "新闻")
		})
	}

	apiRouter := r.Group("/api")
	{
		apiRouter.GET("/", func(context *gin.Context) {
			context.String(200, "我是一个api接口")
		})
		apiRouter.GET("/user", func(context *gin.Context) {
			context.String(200, "userlist的api接口")
		})
	}

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

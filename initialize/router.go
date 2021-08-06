package initialize

import (
	"gin-xutao/middleware"
	"gin-xutao/router"
	"github.com/gin-gonic/gin"
)

func Routers()  *gin.Engine {
	// 1.创建路由
	Router := gin.Default()

	Router.Use(middleware.Cors()) // 如需跨域可以打开
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	//路由分组
	WebRouter := router.RouterGroupApp.WebGroupRouter
	SysRouter := router.RouterGroupApp.SysGroupRouter

	PublicGroup := Router.Group("api")
	PublicGroup.Use(middleware.RestGate())
	{
		SysRouter.InitSystemUserRouter(PublicGroup)
		SysRouter.InitSystemMenuRouter(PublicGroup)

		WebRouter.InitWebUserRouter(PublicGroup)
	}

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	return Router
}
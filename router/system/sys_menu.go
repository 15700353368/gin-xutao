package system

import (
	"gin-xutao/api"
	"github.com/gin-gonic/gin"
)

type SysMenuRouter struct {

}

func (p *SysMenuRouter)InitSystemMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("sys")

	MenuApi := api.ApiGroupApp.SysGroupAPi.SysMenuApi
	{
		baseRouter.GET("/menuList", MenuApi.MenuList)
		baseRouter.GET("/menuInfo", MenuApi.MenuInfo)
		baseRouter.GET("/xutaoTest1", MenuApi.XutaoTest)
	}
	return baseRouter
}
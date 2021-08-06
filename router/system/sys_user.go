package system

import (
	"gin-xutao/api"
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct {

}

//
func (p *SysUserRouter)InitSystemUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("sys")

	UserApi := api.ApiGroupApp.SysGroupAPi.SysUserApi
	{
		baseRouter.POST("/userList", UserApi.UserList)
		baseRouter.GET("/userInfo", UserApi.UserInfo)
		baseRouter.GET("/xutaoTest", UserApi.XutaoTest)
	}

	return baseRouter
}
package web

import (
	"gin-xutao/api"
	"github.com/gin-gonic/gin"
)

type WebUserRouter struct {

}

func (p *WebUserRouter)InitWebUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("web")

	UserApi := api.ApiGroupApp.WebGroupAPi.WebUserApi
	{
		baseRouter.GET("/userList", UserApi.UserList)
		baseRouter.GET("/userInfo", UserApi.UserInfo)
	}

	return baseRouter
}
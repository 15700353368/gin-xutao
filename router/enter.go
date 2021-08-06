package router

import (
	"gin-xutao/router/system"
	"gin-xutao/router/web"
)

type RouterGroup struct {
	SysGroupRouter system.RouterGroup
	WebGroupRouter web.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
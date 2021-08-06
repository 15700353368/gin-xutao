package api

import (
	"gin-xutao/api/system"
	"gin-xutao/api/web"
)

type ApiGroup struct {
	SysGroupAPi system.ApiGroup
	WebGroupAPi web.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
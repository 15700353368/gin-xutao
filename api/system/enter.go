package system

import (
	"gin-xutao/service"
)

type ApiGroup struct {
	SysMenuApi
	SysUserApi
}

var systemUserService = service.ServiceGroupApp.SystemUserService
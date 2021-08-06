package system

import (
	"gin-xutao/global"
	"gin-xutao/model/system"
	paramsSystem "gin-xutao/params/system"
)

type SystemUserService struct {

}

func (p *SystemUserService)GetUserList(paramsSystem.SysUserList) ([]system.SysUsers) {
	var userList []system.SysUsers
	global.GVA_GORM.Find(&userList)
	return userList
}
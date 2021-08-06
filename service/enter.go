package service

import "gin-xutao/service/system"

type ServiceGroup struct {
	SystemUserService system.SystemUserService
}

var ServiceGroupApp ServiceGroup
package system

import (
	"fmt"
	"gin-xutao/params/system"
	"gin-xutao/utils"
	"github.com/gin-gonic/gin"
	"time"
)


type SysUserApi struct {

}

func (p *SysUserApi)UserList(c *gin.Context)  {
	time.Sleep(time.Second*20)

	var sysUser system.SysUserList
	_ = c.ShouldBindJSON(&sysUser)

	data := systemUserService.GetUserList(sysUser)

	//测试redis
	if err := utils.RedisSetString("test","value");err != nil{
		fmt.Println(err.Error())
	}

	c.JSON(200,gin.H{"data":data})
}


func (p *SysUserApi)UserInfo(c *gin.Context)  {
	c.JSON(200,"useInfo：system")
}

func (p *SysUserApi)XutaoTest(c *gin.Context)  {
	c.JSON(200,"menuTest：system")
}
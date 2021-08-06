package system

import "github.com/gin-gonic/gin"


type SysMenuApi struct {

}

func (p *SysMenuApi)MenuList(c *gin.Context)  {
	c.JSON(200,"menuList：system")
}

func (p *SysMenuApi)MenuInfo(c *gin.Context)  {
	c.JSON(200,"menuInfo：system")
}

func (p *SysMenuApi)XutaoTest(c *gin.Context)  {
	c.JSON(200,"menuTest：system")
}
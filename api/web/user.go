package web

import "github.com/gin-gonic/gin"

type WebUserApi struct {

}

func (p *WebUserApi)UserList(c *gin.Context)  {
	c.JSON(200,"useList：web")
}

func (p *WebUserApi)UserInfo(c *gin.Context)  {
	c.JSON(200,"useInfo：web")
}
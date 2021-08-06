package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
package main

import (
	"context"
	"fmt"
	"gin-xutao/global"
	"gin-xutao/initialize"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//yaml文件初始化 到指定struct
	initialize.Configs()

	//初始化日志库zap
	global.GVA_LOG = initialize.Zap()

	//初始化数据库gorm
	global.GVA_GORM = initialize.Gorm()

	//初始化redis
	initialize.Redis()

	//初始化路由
	Router := initialize.Routers()


	//启动服务
	s := initialize.InitServer(global.GVA_CONFIG.System.Addr, Router)

	//默认启动
	//s.ListenAndServe()


	//平滑启动
	gracefulRestart(s)
	//Router.Run(global.GVA_CONFIG.System.Addr)
}


//平滑启动
func gracefulRestart(s *http.Server)  {
	//Initializing the server in a goroutine so that
	//it won't block the graceful shutdown handling below
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.GVA_LOG.Error("listen:", zap.Any("err", err))
		}
	}()

	//创建一个信号监听通道
	quit := make(chan os.Signal, 1)
	//监听 syscall.SIGINT 跟 syscall.SIGTERM信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	si :=<-quit
	fmt.Println("Shutting down server...",si)

	//shutdown方法需要传入一个上下文参数，这里就设计到两种用法
	//1.WithCancel带时间，表示接收到信号之后，过完该断时间不管当前请求是否完成，强制断开
	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	//2.不带时间，表示等待当前请求全部完成再断开
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		//当请求还在的时候强制断开了连接将产生错误，err不为空
		global.GVA_LOG.Error("Server forced to shutdown:", zap.Any("err", err))
	}
	fmt.Println("Server exiting")
}
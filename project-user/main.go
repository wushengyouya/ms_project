package main

import (
	"github.com/gin-gonic/gin"
	srv "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-user/config"
	"github.com/wushengyouya/project-user/router"
)

func main() {
	r := gin.Default()
	config.AppConf.InitZapLog()
	// 注册路由
	router.InitRouter(r)
	// 注册grpc服务
	grpc := router.RegisterGrpc()
	stop := func() {
		grpc.Stop()
	}
	srv.Run(r, "my-project", ":8081", stop)

}

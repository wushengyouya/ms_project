package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-api/config"
	"github.com/wushengyouya/project-api/router"
	srv "github.com/wushengyouya/project-common"
)

func main() {
	r := gin.Default()
	config.AppConf.InitZapLog()
	// 注册路由
	router.InitRouter(r)

	srv.Run(r, config.AppConf.AppConfig.Name, config.AppConf.AppConfig.Addr, nil)

}

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
	router.InitRouter(r)
	srv.Run(r, "my-project", ":8080")

}

package project

import (
	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-api/api/middleware"
)

type RouterProject struct {
}

func init() {}

func (*RouterProject) Register(r *gin.Engine) {
	InitProjectRpcClient()
	h := New()
	group := r.Group("/project")
	group.Use(middleware.TokenVerify())
	group.POST("/index", h.Index)
	group.POST("/selfList", h.MyProjectList)
}

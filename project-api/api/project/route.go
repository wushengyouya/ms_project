package project

import (
	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-api/api/middleware"
)

type RouterProject struct {
}

func init() {}

func (*RouterProject) Register(r *gin.Engine) {
	h := New()
	group := r.Group("/project/index")
	group.Use(middleware.TokenVerify())
	group.POST("", h.Index)
}

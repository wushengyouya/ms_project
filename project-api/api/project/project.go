package project

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	common "github.com/wushengyouya/project-common"
)

type ProjectHandler struct{}

func New() *ProjectHandler {
	return new(ProjectHandler)
}

func (*ProjectHandler) Index(ctx *gin.Context) {
	result := &common.Result{}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 构建IndexMessage 调用 Project rpc服务

}

package project

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-grpc/project"
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
	msg := &project.IndexMessage{}
	indexResp, err := ProjectServiceClient.Index(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(indexResp.Menus))

}

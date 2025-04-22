package project

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/wushengyouya/project-api/pkg/model"
	pro "github.com/wushengyouya/project-api/pkg/model/project"
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

func (*ProjectHandler) MyProjectList(ctx *gin.Context) {
	reuslt := &common.Result{}
	// 获取参数
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	memberIdStr, _ := ctx.Get("memberId")
	memberId := memberIdStr.(int64)
	page := &model.Page{}
	page.Bind(ctx)

	msg := &project.ProjectRpcMessage{MemberId: memberId, Page: page.Page, PageSize: page.PageSize}
	myProjectResponse, err := ProjectServiceClient.FindProjectByMemId(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, reuslt.Fail(code, msg))
		return
	}
	if myProjectResponse.Pm == nil {
		myProjectResponse.Pm = []*project.ProjectMessage{}
	}

	var pms []*pro.ProjectAndMember
	copier.Copy(&pms, myProjectResponse.Pm)
	ctx.JSON(http.StatusOK, reuslt.Success(gin.H{
		"list":  pms,
		"total": myProjectResponse.Total,
	}))
}

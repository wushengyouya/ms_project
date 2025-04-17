package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/wushengyouya/project-api/pkg/model/user"
	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-common/logs"
	"github.com/wushengyouya/project-grpc/user/login"
)

type LoginHanlder struct {
}

func (*LoginHanlder) GetCaptcha(ctx *gin.Context) {

	result := common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 使用grpc客户端调用远程服务
	resp, err := UserClient.GetCaptcha(c, &login.CaptchaMessage{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		logs.LG.Error(err.Error())
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(resp.Code))

}

func (*LoginHanlder) Register(ctx *gin.Context) {
	result := new(common.Result)
	// 1、接收参数
	var req user.RegisterReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}

	// 2、校验参数
	if err := req.Verify(); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
	}
	// 3、调用login rpc服务
	// 加入调用超时
	c, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	msg := &login.RegisterMessage{}
	err = copier.Copy(msg, req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy有误"))
		return
	}
	// 3.使用rpc客户端调用Register服务
	_, err = UserClient.Register(c, msg)
	if err != nil {
		// 转化rpc返回的err
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(""))

}

func New() *LoginHanlder {
	return new(LoginHanlder)
}

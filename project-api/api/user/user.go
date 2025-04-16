package user

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-common/logs"
	loginServiceV1 "github.com/wushengyouya/project-user/pkg/service/login.service.v1"
)

type LoginHanlder struct {
}

func (*LoginHanlder) GetCaptcha(ctx *gin.Context) {

	result := common.Result{}
	mobile := ctx.PostForm("mobile")
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 使用grpc客户端调用远程服务
	resp, err := UserClient.GetCaptcha(c, &loginServiceV1.CaptchaMessage{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		logs.LG.Error(err.Error())
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(resp.Code))

}

func New() *LoginHanlder {
	return new(LoginHanlder)
}

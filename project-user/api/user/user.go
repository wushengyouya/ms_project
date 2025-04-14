package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-user/pkg/dao"
	"github.com/wushengyouya/project-user/pkg/model"
	"github.com/wushengyouya/project-user/pkg/repo"
	"go.uber.org/zap"
)

type HandlerUser struct {
	cache repo.Cache
}

func New() *HandlerUser {
	return &HandlerUser{cache: dao.RC}
}

func (h *HandlerUser) getCaptcha(ctx *gin.Context) {

	resp := &common.Result{}
	// 1.获取参数
	mobile := ctx.PostForm("mobile")
	// 2.校验参数
	if !common.VerifyMobile(mobile) {
		ctx.JSON(http.StatusOK, resp.Fail(model.NoLegalMobile, "手机号不合法"))
		return
	}
	// 3.生成验证码（调用第三平台，放入go协程中执行）
	code := "123456"
	go func() {
		// 模拟
		h.cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		zap.L().Info("短信平台调用成功，发送短信")
		// 4.存储验证码到redis中过期时间15分钟
		zap.L().Info(fmt.Sprintf("将手机号和验证码存入redis成功: REGISTER_%s : %s", mobile, code))

	}()
	ctx.JSON(http.StatusOK, resp.Success("getcaptcha success"))
}

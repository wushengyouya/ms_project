package login_service_v1

import (
	"context"
	"fmt"
	"time"

	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-user/pkg/model"
	"github.com/wushengyouya/project-user/pkg/repo"
	"go.uber.org/zap"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	Cache repo.Cache
}

func (ls *LoginService) GetCaptcha(ctx context.Context, msg *CaptchaMessage) (*CaptchaResponse, error) {
	// 1.获取参数
	mobile := msg.Mobile
	// 2.校验参数
	if !common.VerifyMobile(mobile) {
		return nil, errs.GrpcError(model.NoLegalMobile)
	}
	// 3.生成验证码（调用第三平台，放入go协程中执行）
	code := "123456"
	go func() {
		// 模拟
		err := ls.Cache.Put("REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			zap.L().Info("验证码存入redis发生错误,cause by : ", zap.Error(err))
		}
		zap.L().Info("短信平台调用成功，发送短信")
		zap.L().Debug("短信平台调用成功，发送短信 debug")
		zap.L().Warn("短信平台调用成功，发送短信warn")
		zap.L().Error("短信平台调用成功，发送短信error")
		// 4.存储验证码到redis中过期时间15分钟
		zap.L().Info(fmt.Sprintf("将手机号和验证码存入redis成功: REGISTER_%s : %s", mobile, code))

	}()
	return &CaptchaResponse{Code: int32(2000)}, nil
}

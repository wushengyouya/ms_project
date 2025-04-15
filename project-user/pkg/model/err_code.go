package model

import "github.com/wushengyouya/project-common/errs"

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
	Success       = errs.NewError(2000, "验证码获取成功")
)

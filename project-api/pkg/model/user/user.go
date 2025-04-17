package user

import (
	"errors"

	common "github.com/wushengyouya/project-common"
)

type RegisterReq struct {
	Email     string `json:"email" form:"email"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"passowrd" form:"password"`
	Password2 string `json:"password" form:"passowrd"`
	Mobile    string `json:"mobile" form:"mobile"`
	Captcha   string `json:"captcha" form:"captcha"`
}

func (r RegisterReq) VerifyPassword() bool {
	return r.Password == r.Password2
}

func (r RegisterReq) Verify() error {
	if !common.VerfiyEmailFormat(r.Email) {
		return errors.New("邮箱格式不正确")
	}
	if !common.VerifyMobile(r.Mobile) {
		return errors.New("手机号格式不正确")
	}
	if !r.VerifyPassword() {
		return errors.New("输入的两次密码不一致")
	}
	return nil
}

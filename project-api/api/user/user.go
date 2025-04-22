package user

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/wushengyouya/project-api/pkg/model/user"
	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/encrypts"
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
	log.Println(req)

	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}

	// 2、校验参数
	if err := req.Verify(); err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, err.Error()))
		return
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

func (*LoginHanlder) Login(ctx *gin.Context) {
	// 接收用户名和密码
	result := new(common.Result)
	var req user.LoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "参数格式有误"))
		return
	}
	loginMsg := new(login.LoginMessage)
	err = copier.Copy(loginMsg, req)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy失败"))
		return
	}

	// 调用rpc登录
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	loginMsg.Password = encrypts.Sha256(req.Password)
	loginResp, err := UserClient.Login(c, loginMsg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	resp := new(user.LoginResp)
	err = copier.Copy(resp, loginResp)
	if err != nil {
		ctx.JSON(http.StatusOK, result.Fail(http.StatusBadRequest, "copy失败"))
		return
	}
	ctx.JSON(http.StatusOK, result.Success(resp))
}
func (*LoginHanlder) MyOrgList(ctx *gin.Context) {
	reuslt := common.Result{}
	token := ctx.GetHeader("Authorization")
	// 验证用户是否已经登录
	mem, err := UserClient.TokenVerify(context.Background(), &login.LoginMessage{Token: token})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, reuslt.Fail(code, msg))
		return
	}
	list, err := UserClient.MyOrgList(context.Background(), &login.UserMessage{MemId: mem.Member.Id})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, reuslt.Fail(code, msg))
		return
	}
	if list.OrganizationList == nil {
		ctx.JSON(http.StatusOK, reuslt.Success([]*user.OrganizationList{}))
		return
	}

	var orgs []*user.OrganizationList
	copier.Copy(orgs, list.OrganizationList)
	ctx.JSON(http.StatusOK, reuslt.Success(orgs))
}
func New() *LoginHanlder {
	return new(LoginHanlder)
}

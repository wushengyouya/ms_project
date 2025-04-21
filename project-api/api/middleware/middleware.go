package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-api/api/user"
	common "github.com/wushengyouya/project-common"
	"github.com/wushengyouya/project-common/errs"
	"github.com/wushengyouya/project-grpc/user/login"
)

func TokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 校验token
		result := common.Result{}
		// 1.从header中获取token
		token := ctx.GetHeader("Authorization")
		// 2.调用user服务进行token认证
		c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		resp, err := user.UserClient.TokenVerify(c, &login.LoginMessage{Token: token})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			ctx.JSON(http.StatusOK, result.Fail(code, msg))
			ctx.Abort()
			return
		}
		// 3.处理结果,认证通过，将信息放入gin的上下文，失败则返回未登录
		ctx.Set("memberId", resp.Member.Id)
		ctx.Next()
	}
}

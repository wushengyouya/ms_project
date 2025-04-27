package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-api/pkg/limiter"
	"github.com/wushengyouya/project-api/pkg/model"
	common "github.com/wushengyouya/project-common"
)

type RouterLogin struct {
}

func (rl *RouterLogin) Register(r *gin.Engine) {
	InitUserRpc()

	h := New()
	// 验证码请求增加限流
	r.POST("/project/login/getCaptcha", rateLimiter(), h.GetCaptcha)
	r.POST("/project/login/register", h.Register)
	r.POST("/project/login", h.Login)
	r.POST("/project/organization/_getOrgList", h.MyOrgList)
	r.GET("/ping", h.HealthCheck)
}

// 限流中间件
func rateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化令牌桶
		l := limiter.NewMetHodLimiter().AddBuckets(limiter.LimiterBuckerRule{
			Key:          "/project/login/getCaptcha",
			FillInterval: time.Second,
			Capacity:     10,
			Quantum:      2,
		})
		key := l.Key(ctx)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			// 如果没取到令牌返回err
			if count == 0 {
				ctx.JSON(http.StatusOK, common.Result{Code: model.MoreRequest, Msg: "请求过多"})
				ctx.Abort()
			}
		}
		ctx.Next()
	}
}

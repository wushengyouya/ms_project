package user

import "github.com/gin-gonic/gin"

type RouterLogin struct {
}

func (rl *RouterLogin) Register(r *gin.Engine) {
	InitUserRpc()
	h := New()
	r.POST("/project/login/getCaptcha", h.GetCaptcha)
	r.POST("/project/login/register", h.Register)
	r.POST("/project/login", h.Login)
	r.POST("/project/organization/_getOrgList", h.MyOrgList)
	r.GET("/ping", h.HealthCheck)
}

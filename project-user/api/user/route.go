package user

import "github.com/gin-gonic/gin"

type RouterUser struct{}

func (*RouterUser) Route(r *gin.Engine) {
	handlerUser := New()
	r.POST("/project/login/getCaptcha", handlerUser.getCaptcha)
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-user/api/user"
)

/*
	路由注册
*/

type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct{}

func New() *RegisterRouter {
	return new(RegisterRouter)
}

func (RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Route(r)
}

func InitRouter(r *gin.Engine) {
	router := New()
	router.Route(&user.RouterUser{}, r)

}

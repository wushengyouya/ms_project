package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-api/api/project"
	"github.com/wushengyouya/project-api/api/user"
	"google.golang.org/grpc"
)

/*
	路由注册
*/

type Router interface {
	Register(r *gin.Engine)
}

type RegisterRouter struct{}

func New() *RegisterRouter {
	return new(RegisterRouter)
}

func (RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Register(r)
}

func InitRouter(r *gin.Engine) {
	router := New()
	router.Route(new(user.RouterLogin), r)
	router.Route(new(project.RouterProject), r)
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

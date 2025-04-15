package router

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-user/config"
	"github.com/wushengyouya/project-user/pkg/dao"
	loginServiceV1 "github.com/wushengyouya/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
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

}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.AppConf.GrpcConfig.Addr,
		RegisterFunc: func(g *grpc.Server) {
			loginServiceV1.RegisterLoginServiceServer(g, &loginServiceV1.LoginService{Cache: dao.RC})
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", config.AppConf.GrpcConfig.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Println("server started error: ", err)
			return
		}
	}()
	return s
}

package router

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/wushengyouya/project-common/discovery"
	"github.com/wushengyouya/project-common/logs"
	"github.com/wushengyouya/project-user/config"
	loginServiceV1 "github.com/wushengyouya/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
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
			loginServiceV1.RegisterLoginServiceServer(g, loginServiceV1.New())
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
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

func RegisterEtcdServer() {
	// 1. 创建一个基于 etcd 的服务发现解析器（Resolver）
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs, logs.LG)
	// 2. 将该解析器注册到 gRPC 的解析器注册表中
	/*
		将自定义的 Resolver 注册到 gRPC 的全局解析器注册表中。
		此后，当 gRPC 客户端连接目标服务时（如 dial("etcd:///my-service")），
		会通过该 Resolver 从 etcd 获取实际地址。
	*/
	resolver.Register(etcdRegister)

	info := discovery.Server{
		Name:    config.AppConf.GrpcConfig.Name,
		Addr:    config.AppConf.GrpcConfig.Addr,
		Version: config.AppConf.GrpcConfig.Version,
		Weight:  config.AppConf.GrpcConfig.Weight,
	}

	r := discovery.NewRegister(config.AppConf.EtcdConfig.Addrs, logs.LG)

	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatalln(err)
	}

}

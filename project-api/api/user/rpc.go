package user

import (
	"log"

	"github.com/wushengyouya/project-api/config"
	"github.com/wushengyouya/project-common/discovery"
	"github.com/wushengyouya/project-common/logs"
	loginServiceV1 "github.com/wushengyouya/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var UserClient loginServiceV1.LoginServiceClient

func InitUserRpc() {
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.NewClient(etcdRegister.Scheme()+":///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connet：%v", err)
	}
	UserClient = loginServiceV1.NewLoginServiceClient(conn)
}

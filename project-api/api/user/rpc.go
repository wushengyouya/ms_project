package user

import (
	"log"

	"github.com/wushengyouya/project-api/config"
	"github.com/wushengyouya/project-common/discovery"
	"github.com/wushengyouya/project-common/logs"
	"github.com/wushengyouya/project-grpc/user/login"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var UserClient login.LoginServiceClient

func InitUserRpc() {
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.NewClient(etcdRegister.Scheme()+":///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connet：%v", err)
	}
	UserClient = login.NewLoginServiceClient(conn)
}

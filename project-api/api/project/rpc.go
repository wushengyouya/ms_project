package project

import (
	"log"

	"github.com/wushengyouya/project-api/config"
	"github.com/wushengyouya/project-common/discovery"
	"github.com/wushengyouya/project-common/logs"
	"github.com/wushengyouya/project-grpc/project"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var ProjectServiceClient project.ProjectServiceClient

func InitProjectRpcClient() {
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)

	conn, err := grpc.NewClient(etcdRegister.Scheme()+":///project", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connet: %v", err)
	}
	ProjectServiceClient = project.NewProjectServiceClient(conn)

}

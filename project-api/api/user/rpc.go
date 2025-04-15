package user

import (
	"log"

	loginServiceV1 "github.com/wushengyouya/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var UserClient loginServiceV1.LoginServiceClient

func InitUserRpc() {
	conn, err := grpc.NewClient(":8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connet : %v", err)

	}
	UserClient = loginServiceV1.NewLoginServiceClient(conn)
}

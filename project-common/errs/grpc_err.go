package errs

import (
	common "github.com/wushengyouya/project-common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcError(e *BError) error {
	return status.Error(codes.Code(e.Code), e.Msg)
}

func ParseGrpcError(err error) (common.BusinessCode, string) {
	s, _ := status.FromError(err)
	return common.BusinessCode(s.Code()), s.Message()
}

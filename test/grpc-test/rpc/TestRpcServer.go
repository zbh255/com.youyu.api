package rpc

import (
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/test/grpc-test/rpc/pb"
	"context"
)

type TestServerAPi struct {
	pb.UnimplementedTestApiServer
}

func (t *TestServerAPi) TestInt(ctx context.Context, data *pb.TestData) (*pb.TestData, error) {
	return data,status.Error(ecode.UserDuplicate,"没有该用户")
}


package client

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"github.com/rs/zerolog/log"
)

// IOW 实现io.Write接口，用于向grpc server传递数据
type IOW struct {
	CentRpcPushStream rpc.CentApi_PushLogStreamClient
	FileName          string
}

func (i *IOW) Write(p []byte) (int, error) {
	err := i.CentRpcPushStream.Send(&rpc.Log{
		FileName: i.FileName,
		Value:    p,
	})
	return len(p), err
}

// 测试方法
type IOWTEST struct {
}

func (i *IOWTEST) Write(p []byte) (int, error) {
	log.Print(string(p))
	return 0, nil
}

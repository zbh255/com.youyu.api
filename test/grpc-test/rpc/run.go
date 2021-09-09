package rpc

import (
	"com.youyu.api/test/grpc-test/rpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
)

func Run() {
	listener, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer( // grpc超时连接设置
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    10,
			Timeout: 3,
		}))
	pb.RegisterTestApiServer(grpcServer, &TestServerAPi{})
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

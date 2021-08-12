package main

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/app/rpc/server"
	"com.youyu.api/common/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	var business config.Config = &config.BusinessConfig{}
	businessConf,err := business.GetConfig()
	// TODO:优化错误处理
	if err != nil {
		panic(err)
	}
	result := businessConf.(*config.BusinessConfig)
	listener, err := net.Listen("tcp", result.DataRPCServer.IP + ":" + result.DataRPCServer.Port)
	if err != nil {
		log.Fatal("cannot create a listener at the address")
	}
	grpcServer := grpc.NewServer()
	rpc.RegisterMysqlApiServer(grpcServer, &server.MysqlApiServer{})
	log.Fatalln(grpcServer.Serve(listener))
}

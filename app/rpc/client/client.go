package client

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/common/config"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
)

func GetMysqlApiRpcServerLink(result *config.BusinessConfig) (rpc.MysqlApiClient, *grpc.ClientConn) {
	conn, err := grpc.Dial(result.DataRPCServer.IP+":"+result.DataRPCServer.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc business_server")
	}
	return rpc.NewMysqlApiClient(conn), conn
}

func GetCentApiRpcServerLink(result *config.BusinessConfig) (rpc.CentApiClient,*grpc.ClientConn,error)  {
	conn, err := grpc.Dial(result.CentRPCServer.IP+":"+result.CentRPCServer.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil,nil,errors.Wrap(errors.New("client cannot dial grpc business_server"),"")
	}
	return rpc.NewCentApiClient(conn), conn,nil
}
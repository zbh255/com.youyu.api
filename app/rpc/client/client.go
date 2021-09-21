// 本包存放各种获取业务相关的Rpc call的原语
package client

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/config"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)


func GetCentApiRpcServerLink(result *config.BusinessConfig) (rpc.CentApiClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(result.CentRPCServer.IP+":"+result.CentRPCServer.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, errors.Wrap(errors.New("client cannot dial grpc cent_rpc_server"), "")
	}
	return rpc.NewCentApiClient(conn), conn, nil
}


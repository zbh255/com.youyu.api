package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/config"
	"github.com/pkg/errors"
	"github.com/silenceper/pool"
	"google.golang.org/grpc"
)

// 该文件将grpc的连接和business用到的配置存放统一

type ConnectAndConfig struct {
	Config               *config.AutoGenerated
	DataRpcConnPool      pool.Pool
	SecretKeyRpcConnPool pool.Pool
}

var ConnectAndConf *ConnectAndConfig

// GetDataRpcServer 向连接池取得连接
func GetDataRpcServer(meta interface{}, err error) (rpc.MysqlApiClient, *grpc.ClientConn, error) {
	if err != nil {
		return nil, nil, errors.Wrap(err, "get grpc link failed")
	} else {
		m := meta.(*[2]interface{})
		return m[0].(rpc.MysqlApiClient), m[1].(*grpc.ClientConn), nil
	}
}

// 向连接池取得token和密钥rpc的连接
func GetSecretKeyRpcServer(meta interface{}, err error) (rpc.SecretKeyApiClient, *grpc.ClientConn, error) {
	if err != nil {
		return nil, nil, errors.Wrap(err, "get grpc link failed")
	} else {
		m := meta.(*[2]interface{})
		return m[0].(rpc.SecretKeyApiClient), m[1].(*grpc.ClientConn), nil
	}
}

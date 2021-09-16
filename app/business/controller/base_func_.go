// 抽象出来的公共函数
package controller

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LinkDataRpc() (interface{}, rpc.MysqlApiClient, error) {
	lis, err := ConnectAndConf.DataRpcConnPool.Get()
	dataClient, _, err := GetDataRpcServer(lis, err)
	if err != nil {
		return nil, nil, err
	}
	return lis, dataClient, err
}

func LinkSecretKeyRpc() (interface{}, rpc.SecretKeyApiClient,error)  {
	// 连接secretKey_rpc
	secretKeyLis, err := ConnectAndConf.SecretKeyRpcConnPool.Get()
	secretKeyClient, _, err := GetSecretKeyRpcServer(secretKeyLis, err)
	if err != nil {
		return nil, nil, err
	}
	return secretKeyLis,secretKeyClient,nil
}

func ReturnServerErrJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.ServerErr.Code(),
		"message": ecode.ServerErr.Message(),
	})
}

func ReturnJsonParseErrJson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    ecode.JsonParseError.Code(),
		"message": ecode.JsonParseError.Message(),
	})
}

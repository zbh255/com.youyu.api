package main

import (
	"com.youyu.api/app/business/controller"
	"com.youyu.api/app/rpc/client"
	"com.youyu.api/app/rpc/data/model"
	"com.youyu.api/app/rpc/data/option"
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/config"
	"com.youyu.api/lib/database"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/path"
	"com.youyu.api/lib/router"
	"context"
	"encoding/json"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	zlg "github.com/rs/zerolog/log"
	"io"
	"os"
	"time"
)

func main() {
	// 初始化配置
	// panic均为应用不能正常运行的情况
	var business config.Config = &config.BusinessConfig{}
	businessConf, err := business.GetConfig()
	if err != nil {
		panic(err)
	}
	result := businessConf.(*config.BusinessConfig)
	// 连接配置中心获取配置
	clientE, conn, err := client.GetCentApiRpcServerLink(result)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	resultByte, err := clientE.GetRpcServerConfFile(context.Background(), &rpc.Null{})
	if err != nil {
		panic(err)
	}
	generated := config.AutoGenerated{}
	resultI, err := generated.Unmarshal(resultByte.Data)
	if err != nil {
		panic(err)
	}
	resultAG := resultI.(*config.AutoGenerated)
	stream, err := clientE.PushLogStream(context.Background())
	if err != nil {
		panic(err)
	}
	defer stream.CloseSend()
	r := gin.New()
	gin.DefaultWriter = io.MultiWriter(&client.IOW{
		CentRpcPushStream: stream,
		FileName:          path.LogWebServerFileName,
	}, os.Stdout)
	// json日志
	r.Use(logger.SetLogger())
	// 初始化业务日志
	businessStream, err := clientE.PushLogStream(context.Background())
	if err != nil {
		panic(err)
	}
	defer businessStream.CloseSend()
	router.InitRouter(r, log.Logger(&log.ZLogger{
		Level: zerolog.ErrorLevel,
		Logger: zlg.Output(&client.IOW{
			CentRpcPushStream: businessStream,
			FileName:          path.LogBusinessFileName,
		}),
	}))
	//// 初始化grpc设置
	//grpc.MaxSendMsgSize(2 << 31)
	//grpc.MaxRecvMsgSize(2 << 31)
	//grpc.InitialWindowSize(2 << 29)
	//grpc.InitialConnWindowSize(2 << 29)
	//grpc.MaxConcurrentStreams(2 << 8)
	// 初始化redis连接池客户端
	// 向option注册连接池
	redisPool := redis.NewClient(&redis.Options{
		// 网络连接方式
		Network: "tcp",
		// redis server地址
		Addr: resultAG.Redis.IPAddr + ":" + resultAG.Redis.Port,
		// redis的密码
		Password: resultAG.Redis.Password,
		// 使用的数据库编号
		DB: 0,
		// 客户端建立连接的超时时间
		DialTimeout: time.Duration(resultAG.Redis.Sync.DialTimeout) * time.Second,
		// 连接池的最大闲置连接数
		PoolSize: resultAG.Redis.Sync.MaxOpenConnSize,
		// 连接池的最小保持的限制连接数量
		MinIdleConns: resultAG.Redis.Sync.MinOpenConnSize,
		// 连接池连接的保活时间
		MaxConnAge: time.Duration(resultAG.Redis.Sync.MaxConnLifeTime) * time.Second,
		// 当连接池中没有空闲连接时，程序等待空闲连接的最长时间
		PoolTimeout: time.Duration(resultAG.Redis.Sync.PoolTimeout) * time.Second,
		// 闲置连接的超时时间
		IdleTimeout: time.Duration(resultAG.Redis.Sync.IdleTimeout) * time.Second,
	})
	if err != nil {
		panic(err)
	}
	// 初始化mysql数据库
	dbInterface := database.DataBase(&database.Mysql{})
	dbInterface.SetConfig(resultAG)
	db,err := dbInterface.GetConnect()
	if err != nil {
		panic(err)
	}
	model.DB = db
	// 初始化给business使用的全局接口，并初始化各自模块的日志
	// database日志
	dbStream, err := clientE.PushLogStream(context.Background())
	if err != nil {
		panic(err)
	}
	defer dbStream.CloseSend()
	// secretKey日志
	secretKeyStream,err := clientE.PushLogStream(context.Background())
	if err != nil {
		panic(err)
	}
	defer secretKeyStream.CloseSend()
	controller.ConnectAndConf = &controller.ConnectAndConfig{
		Config:               resultAG,
		SecretKeyLink: &option.SecretKeyApiServer{
			RedisClient:                redisPool,
			Logger:                     log.Logger(&log.ZLogger{
				Level:  zerolog.ErrorLevel,
				Logger: zlg.Output(&client.IOW{
					CentRpcPushStream: secretKeyStream,
					FileName:          path.LogSecretRpcFileName,
				}),
			}),
		},
		DataBaseLink: &option.MysqlApiServer{
			Logger:                      log.Logger(&log.ZLogger{
				Level:  zerolog.ErrorLevel,
				Logger: zlg.Output(&client.IOW{
					CentRpcPushStream: dbStream,
					FileName:          path.LogDataRpcFileName,
				}),
			}),
		},
	}
	// 注册签名密钥
	controller.TokenSigningKey = []byte(resultAG.Project.Auth.TokenSigntureKey)
	// 注册签钥服务器的密钥
	option.TokenSigntureKey = []byte(resultAG.Project.Auth.TokenSigntureKey)
	// 注册Token过期时间
	controller.TokenExpTime = resultAG.Project.Auth.TokenTimeout
	// 注册错误
	errInfo, err := clientE.GetErrMsgJsonBytes(context.Background(), &rpc.Null{})
	if err != nil {
		panic(err)
	}
	errCodeMap := make(map[int]string)
	err = json.Unmarshal(errInfo.Data, &errCodeMap)
	if err != nil {
		panic(err)
	}
	ecode.Register(errCodeMap)
	_ = r.Run(resultAG.Server.IPAddr + ":" + resultAG.Server.Port)
}

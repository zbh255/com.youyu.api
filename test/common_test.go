package test

import (
	"com.youyu.api/lib/config"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"strconv"
	"testing"
)

func TestConfig(t *testing.T) {
	// business and data_rpc 配置文件
	var business config.Config = &config.BusinessConfig{}
	businessConf, err := business.GetConfig()
	// TODO:优化错误处理
	result := businessConf.(*config.BusinessConfig)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result)
	}
	// 序列化和反序列化
	byteData := result.Marshal()
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log(string(byteData))
	}
	newResult, err := result.Unmarshal(byteData)
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log(newResult.(*config.BusinessConfig))
	}
	// app配置文件
	app := config.Config(&config.AutoGenerated{})
	appConf, err := app.GetConfig()
	resultApp := appConf.(*config.AutoGenerated)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(resultApp)
	}
	// 序列化和反序列化
	byteData = resultApp.Marshal()
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log(string(byteData))
	}
	newResult, err = resultApp.Unmarshal(byteData)
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log(newResult.(*config.AutoGenerated))
	}
}

// TODO GRPC连接池已经废弃
func TestGrpcPool(t *testing.T) {
	//// 读取配置文件
	//resultApp := config.Config(&config.AutoGenerated{})
	//result, err := resultApp.GetConfig()
	//if err != nil {
	//	t.Error(err)
	//}
	//resultAG := result.(*config.AutoGenerated)
	//// 初始化grpc设置
	//grpc.MaxSendMsgSize(2 << 31)
	//grpc.MaxRecvMsgSize(2 << 31)
	//grpc.InitialWindowSize(2 << 29)
	//grpc.InitialConnWindowSize(2 << 29)
	//grpc.MaxConcurrentStreams(2 << 8)
	//// 初始化连接池
	//Factory := func() (interface{}, error) {
	//	m, c, e := client.GetMysqlApiRpcServerLink(resultAG)
	//	return &[2]interface{}{m, c}, e
	//}
	//Close := func(i interface{}) error {
	//	i2 := i.(*[2]interface{})
	//	conn := i2[1].(*grpc.ClientConn)
	//	return conn.Close()
	//}
	//p, err := pool.NewChannelPool(&pool.Config{
	//	InitialCap:  5,
	//	MaxCap:      30,
	//	MaxIdle:     20,
	//	Factory:     Factory,
	//	Close:       Close,
	//	IdleTimeout: 15 * time.Second,
	//})
	//if err != nil {
	//	fmt.Println("err=", err)
	//}
	//controller.ConnectAndConf = &controller.ConnectAndConfig{
	//	Config:          resultAG,
	//	DataRpcConnPool: p,
	//}
	//meta, err := controller.ConnectAndConf.DataRpcConnPool.Get()
	//_, _, err = controller.GetDataRpcServer(meta, err)
	//if err != nil {
	//	t.Error(err)
	//}
	//err = controller.ConnectAndConf.DataRpcConnPool.Put(meta)
	//if err != nil {
	//	t.Error(err)
	//}
}

func TestCosSignature(t *testing.T) {
	conf,err := config.Config(&config.AutoGenerated{}).GetConfig()
	if err != nil {
		t.Error(err)
	}
	result := conf.(*config.AutoGenerated)
	// 新建客户端
	c := sts.NewClient(result.Project.Cos.SecretID,result.Project.Cos.SecretKey,nil)
	opt := &sts.CredentialOptions{
		Policy:          &sts.CredentialPolicy{
			Version:   "2.0",
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + result.Project.Cos.PublicSourceBucket.Region + ":uid/" +
							strconv.Itoa(result.Project.Cos.Appid) + ":" + result.Project.Cos.PublicSourceBucket.Name + "/images/xx.jpg",
					},
				},
			},
		},
		Region:          result.Project.Cos.PublicSourceBucket.Region,
		DurationSeconds: int64(result.Project.Cos.DurationSeconds),
	}
	credential, err := c.GetCredential(opt)
	if err != nil {
		t.Error(err)
	}
	t.Log(credential)
	t.Log(credential.Credentials.TmpSecretID)
	t.Log(credential.Credentials.TmpSecretKey)
	t.Log(credential.Credentials.SessionToken)
}

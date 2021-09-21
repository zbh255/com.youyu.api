package test

import (
	"com.youyu.api/app/business/controller"
	rpcClient "com.youyu.api/app/rpc/client"
	"com.youyu.api/app/rpc/data/model"
	"com.youyu.api/app/rpc/data/option"
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/config"
	"com.youyu.api/lib/database"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/path"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
	zlg "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func InitTestENV() {
	// 初始化配置
	// panic均为应用不能正常运行的情况
	var business config.Config = &config.BusinessConfig{}
	businessConf, err := business.GetConfig()
	if err != nil {
		panic(err)
	}
	result := businessConf.(*config.BusinessConfig)
	// 连接配置中心获取配置
	clientE, conn, err := rpcClient.GetCentApiRpcServerLink(result)
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
				Logger: zlg.Output(&rpcClient.IOW{
					CentRpcPushStream: secretKeyStream,
					FileName:          path.LogSecretRpcFileName,
				}),
			}),
		},
		DataBaseLink: &option.MysqlApiServer{
			Logger:                      log.Logger(&log.ZLogger{
				Level:  zerolog.ErrorLevel,
				Logger: zlg.Output(&rpcClient.IOW{
					CentRpcPushStream: dbStream,
					FileName:          path.LogDataRpcFileName,
				}),
			}),
		},
	}
	// 注册签名密钥
	controller.TokenSigningKey = []byte(resultAG.Project.Auth.TokenSigntureKey)
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
}

// 测试article数据库模型，包含子表
func TestMysqlApiArticle(t *testing.T) {
	InitTestENV()
	client := controller.TakeDataBaseLink()
	article, err := client.AddArticle(context.Background(), &rpc.Article{
		ArticleAbstract:   "hello worlds",
		ArticleContent:    "hello world,hello world",
		ArticleTitle:      "sb",
		ArticleTag:        []string{"11", "linux"},
		Uid:               3,
		ArticleCreateTime: time.Now().Unix(),
		ArticleUpdateTime: time.Now().Unix(),
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test addArticle ok")
	}
	_, err = client.AddArticleStatisticsHot(context.Background(), &rpc.ArticleRequest{ArticleId: []string{article.GetArticleId()}})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test addArticleStatisticsHot ok")
	}

	_, err = client.AddArticleStatisticsFabulous(context.Background(), &rpc.ArticleRequestOne{ArticleId: article.GetArticleId()})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test addArticleStatisticsFabulous ok")
	}

	_, err = client.AddArticleStatisticsCommentNum(context.Background(), &rpc.ArticleRequest{ArticleId: []string{article.GetArticleId()}})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test addArticleStatisticsCommentNum ok")
	}

	article, err = client.GetArticle(context.Background(), &rpc.ArticleRequest{ArticleId: []string{article.GetArticleId()}})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(article)
		t.Log("test getArticle ok")
	}
	time.Sleep(time.Second * 2)
	_, err = client.UpdateArticle(context.Background(), &rpc.Article{
		ArticleId:         article.GetArticleId(),
		ArticleAbstract:   "世界你好",
		ArticleContent:    "我是世界",
		ArticleTitle:      "世界是我",
		ArticleTag:        []string{"linux", "macos"},
		Uid:               10,
		ArticleCreateTime: 0,
		ArticleUpdateTime: 0,
	})
	// 解包错误
	st, _ := status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		result, _ := client.GetArticle(context.Background(), &rpc.ArticleRequest{ArticleId: []string{article.GetArticleId()}})
		t.Log(result)
		t.Log("test updateArticle ok" + st.Message)
	}
	// 连接查询
	result, err := client.GetArticleList(context.Background(), &rpc.ArticleOptions{
		Type:    "desc",
		Order:   "hot",
		Page:    1,
		PageNum: 3,
	})
	st,_ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log(result)
		t.Log("test getArticleList ok" + st.Message)
	}
	_, err = client.DelArticle(context.Background(), &rpc.ArticleRequest{ArticleId: []string{article.GetArticleId()}})
	st, _ = status.FromError(err)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test delArticle ok" + st.Message)
	}
}

// 测试advertisement数据库模型
func TestMysqlApiAdvertisement(t *testing.T) {
	InitTestENV()
	client := controller.TakeDataBaseLink()
	if client == nil {
		panic("client is nil pointer ")
	}
	testData := &rpc.Advertisement{
		AdvertisementId:     1,
		AdvertisementType:   2,
		AdvertisementLink:   "https://xiao-hui.net/NewMysqlApiClient",
		AdvertisementWeight: 9,
		AdvertisementBody:   "https://tencent/video/11",
		AdvertisementOwner:  "youyu.Inc",
	}
	_, err := client.AddAdvertisement(context.Background(), testData)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test addAdvertisement ok")
	}

	result, err := client.GetAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: testData.AdvertisementId})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result)
		t.Log("test getAdvertisement ok")
	}

	// update
	testData.AdvertisementOwner = "8086 Inc"
	testData.AdvertisementBody = "https://tencent/video/13"
	_, err = client.UpdateAdvertisement(context.Background(), testData)
	if err != nil {
		t.Error(err.Error())
	} else {
		result, _ := client.GetAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: testData.AdvertisementId})
		t.Log(result)
		t.Log("test updateAdvertisement ok")
	}
	// 测试广告列表
	testData.AdvertisementId = 2
	_, _ = client.AddAdvertisement(context.Background(), testData)
	testData.AdvertisementId = 3
	_, _ = client.AddAdvertisement(context.Background(), testData)
	results, err := client.GetAdvertisementList(context.Background(), &rpc.ArticleOptions{
		Type:    "desc",
		Order:   "",
		Page:    1,
		PageNum: 3,
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(results)
		t.Log("test getAdvertisements ok")
	}
	// del
	_, err = client.DelAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: testData.AdvertisementId})
	_, err = client.DelAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: 2})
	_, err = client.DelAdvertisement(context.Background(), &rpc.AdvertisementRequest{AdvertisementId: 1})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result)
		t.Log("test delAdvertisement ok")
	}
}

func TestMysqlTags(t *testing.T) {
	InitTestENV()
	client := controller.TakeDataBaseLink()
	if client == nil {
		panic("client is nil pointer")
	}
	// 测试模型
	rand.Seed(time.Now().UnixNano())
	// 限定text长度为10
	// 构建400个text
	defaultTexts := make([]string, 400)
	for k := range defaultTexts {
		defaultTexts[k] = strconv.FormatInt(rand.Int63n(999999999), 10)
	}
	//defaultText := strconv.FormatInt(rand.Int63n(999999999), 10)
	err := error(nil)
	for _, v := range defaultTexts {
		_, err = client.AddTag(context.Background(), &rpc.Tag{Text: []string{v}})
	}
	st, _ := status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("test add tag ok")
	}
	tag2, err := client.GetTagInt32Id(context.Background(), &rpc.Tag{Text: defaultTexts})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("test get tag int32 id ok")
		t.Log(tag2.Tid)
	}
	tag, err := client.GetTagText(context.Background(), &rpc.Tag{Tid: tag2.Tid})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("test get tag Text ok")
		t.Log(tag.Text)
	}
	for _, v := range tag2.Tid {
		_, err = client.DelTag(context.Background(), &rpc.Tag{Tid: []int32{v}})
	}
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("test del tag ok")
	}
}

func TestCentApi(t *testing.T) {
	var busiess config.Config = &config.BusinessConfig{}
	r, err := busiess.GetConfig()
	if err != nil {
		t.Error("client cannot dial grpc business_server")
	}
	result := r.(*config.BusinessConfig)
	conn, err := grpc.Dial(result.CentRPCServer.IP+":"+result.CentRPCServer.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Error("client cannot dial grpc business_server")
	}
	defer conn.Close()
	client := rpc.NewCentApiClient(conn)
	conf, err := client.GetRpcServerConfFile(context.Background(), &rpc.Null{})
	if err != nil {
		t.Error(err)
	}
	t.Log(string(conf.Data))
	_, err = client.SetRpcServerConfFile(context.Background(), &rpc.Config{
		Type: path.ConfRpcRequestType,
		Data: conf.Data,
	})
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log("test SetRpcServerConfFile ok")
	}
	// business
	conf, err = client.GetBusinessConfFile(context.Background(), &rpc.Null{})
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log("test SetRpcServerConfFile ok")
	}
	businessConfig := config.BusinessConfig{}
	newResult, err := businessConfig.Unmarshal(conf.Data)
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log(newResult)
		t.Log("test getBusinessConfFile ok")
	}
	_, err = client.SetBusinessConfFile(context.Background(), &rpc.Config{
		Type: path.ConfBusinessRequestType,
		Data: conf.Data,
	})
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log("test setBusinessConfFile ok")
	}
	client.PushLogStream(context.Background())
	// err_msg的并发设置和获取
	info, err := client.GetErrMsgJsonBytes(context.Background(), &rpc.Null{})
	if err != nil {
		t.Errorf("%+v", err)
	} else {
		t.Log("test getErr_msg ok")
	}
	tmpMap := make(map[int]string)
	err = json.Unmarshal(info.Data, &tmpMap)
	if err != nil {
		t.Error(err)
	}
	tmpMap[-404] = "页面飞走了"
	info.Data, err = json.Marshal(tmpMap)
	if err != nil {
		t.Errorf("%+v", err)
	}
	_, err = client.SetErrMsgJson(context.Background(), &rpc.Info{Type: path.ErrMsgJsonFileName, Data: info.Data})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test setErr_msg ok")
	}
}

func TestSecretKeyApi(t *testing.T) {
	InitTestENV()
	client := controller.TakeSecretKeyLink()
	rand.Seed(time.Now().UnixNano())
	uid := rand.Int31()
	user, err := client.BindTokenToUser(context.Background(), &rpc.User{Uid: uid, ExpTime: int64(time.Second * 100)})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test bind token to user ok")
		t.Log(user)
	}
	user, err = client.ForTokenGetBindUser(context.Background(), &rpc.User{Token: user.Token})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test for token get bind user ok")
		t.Log(user)
	}
	_, err = client.DeleteBindUser(context.Background(), user)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test deleted bind user ok")
	}
	// 通过client id来设置和获取
	clientId := "1"
	key, err := client.GetPublicKey(context.Background(), &rpc.RsaKey{ClientId: clientId})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(key.PublicKey)
	}
	key, err = client.GetPrivateKey(context.Background(), &rpc.RsaKey{ClientId: clientId})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(key.PrivateKey)
	}
}

func TestMysqlUserLoginAndSign(t *testing.T) {
	InitTestENV()
	client := controller.TakeDataBaseLink()
	// 测试模型
	userName := "xiao-hui-xx"
	userPassword := "womeiyoumima"
	_, err := client.CreateUserSign(context.Background(), &rpc.UserSign{
		UserName:     userName,
		UserPassword: userPassword,
		UserBindInfo: "12345678901",
		VCode:        "1234",
		VToken:       "12345",
		SignType:     rpc.LoginAndSignType_Native,
	})
	st, _ := status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("create user sign ok")
	}
	// 重试构造错误
	_, err = client.CreateUserSign(context.Background(), &rpc.UserSign{
		UserName:     userName,
		UserPassword: userPassword,
		UserBindInfo: "12345678901",
		VCode:        "1234",
		VToken:       "12345",
		SignType:     rpc.LoginAndSignType_Native,
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Log("create user sign err ok")
	} else {
		t.Error(st.Code)
		t.Error(st.Message)
	}
	// 测试验证用户
	baseData, err := client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		UserName:     userName,
		UserPassword: userPassword,
		Save:         0,
		LoginType:    rpc.LoginAndSignType_Native,
		WechatData:   nil,
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("check user ok")
		t.Log(baseData.Data["uid"])
		t.Log(baseData.Data["user_name"])
	}
	_, err = client.DeleteUserSign(context.Background(), &rpc.UserAuth{Uid: baseData.Data["uid"]})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("del user sign ok")
	}
	// 测试其他注册登录方式
	// 测试通过手机号码注册用户
	userName_phoneTest := "xiao-hui_ph"
	userPassword_phoneTest := "16878q37q25"
	_, err = client.CreateUserSign(context.Background(), &rpc.UserSign{
		UserName:     userName_phoneTest,
		UserPassword: userPassword_phoneTest,
		UserBindInfo: "13025800995",
		SignType:     rpc.LoginAndSignType_Phone,
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("create phone user sign ok")
	}
	// 通过手机号码验证用户
	data, err := client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		UserBindInfo: "13025800995",
		Save:         0,
		LoginType:    rpc.LoginAndSignType_Phone,
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
		panic("")
	} else {
		t.Log("check phone user ok")
		t.Log(data.Data["user_name"])
		t.Log(data.Data["uid"])
	}
	// 删除用户
	_, err = client.DeleteUserSign(context.Background(), &rpc.UserAuth{Uid: data.Data["uid"]})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("create phone user sign ok")
	}
	// 测试其他注册登录方式
	// 测试通过微信注册用户
	userName_wechatTest := "xiao-hui_we"
	userPassword_wechatTest := "16878q37q25"
	wechat_openid := "20060201199"
	_, err = client.CreateUserSign(context.Background(), &rpc.UserSign{
		UserName:     userName_wechatTest,
		UserPassword: userPassword_wechatTest,
		SignType:     rpc.LoginAndSignType_Wechat,
		WechatData: &rpc.WechatUserinfo{
			NickName:  "小辉",
			AvatarUrl: "https://helloworld.com/url",
			Gender:    2,
			Country:   "中国",
			Province:  "广东",
			City:      "江门",
			Language:  "zh_CN",
			Openid:    wechat_openid,
		},
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("create wechat user sign ok")
	}
	// 通过微信openid验证用户
	baseData, err = client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		Save:       0,
		LoginType:  rpc.LoginAndSignType_Wechat,
		WechatData: &rpc.WechatUserinfo{Openid: wechat_openid},
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("check wechat user ok")
		t.Log(data.Data["user_name"])
		t.Log(data.Data["uid"])
	}
	// 删除用户
	_, err = client.DeleteUserSign(context.Background(), &rpc.UserAuth{Uid: baseData.Data["uid"]})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("del wechat user sign ok")
	}
	// 测试原生方式注册用户并添加验证方式
	userName = "xiao-hui-xx_native"
	userPassword = "womeiyoumima"
	_, err = client.CreateUserSign(context.Background(), &rpc.UserSign{
		UserName:     userName,
		UserPassword: userPassword,
		SignType:     rpc.LoginAndSignType_Native,
	})
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("create user sign ok")
	}
	baseData, err = client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		UserName:     userName,
		UserPassword: userPassword,
		Save:         0,
		LoginType:    rpc.LoginAndSignType_Native,
		WechatData:   nil,
	})
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("check native user ok")
	}
	// 添加邮箱验证方式
	email := "565574327@qq.com"
	_, err = client.AddUserCheckInfoEmail(context.Background(), &rpc.UserCheckEmail{
		Email: email,
		Code:  637981,
		Ua:    &rpc.UserAuth{Uid: baseData.Data["uid"]},
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("add check inf email ok")
	}
	// 添加手机验证方式
	phone := 13025801998
	_, err = client.AddUserCheckInfoPhone(context.Background(), &rpc.UserCheckPhone{
		Phone: int64(phone),
		Code:  657890,
		Ua:    &rpc.UserAuth{Uid: baseData.Data["uid"]},
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("del wechat user sign ok")
	}
	// 添加微信验证方式
	openid := "wexinyonghu3389"
	_, err = client.AddUserCheckInfoWechat(context.Background(), &rpc.UserCheckWechat{
		Openid: openid,
		Code:   "898989",
		Ua:     &rpc.UserAuth{Uid: baseData.Data["uid"]},
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("del wechat user sign ok")
	}
	// 验证手机号登录
	newBaseData, err := client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		UserBindInfo: strconv.Itoa(phone),
		VCode:        "657890",
		Save:         0,
		LoginType:    rpc.LoginAndSignType_Phone,
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		if newBaseData.Data["uid"] == baseData.Data["uid"] {
			t.Log("check add phone check info ok")
		}
	}
	// 验证邮箱登录
	newBaseData, err = client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		UserBindInfo: email,
		VCode:        "657890",
		Save:         0,
		LoginType:    rpc.LoginAndSignType_Email,
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		if newBaseData.Data["uid"] == baseData.Data["uid"] {
			t.Log("check add email check info ok")
		}
	}
	// 验证微信登录
	newBaseData, err = client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		VCode:      "657890",
		Save:       0,
		LoginType:  rpc.LoginAndSignType_Wechat,
		WechatData: &rpc.WechatUserinfo{Openid: openid},
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		if newBaseData.Data["uid"] == baseData.Data["uid"] {
			t.Log("check add wechat openid check info ok")
		}
	}
	// 删除添加的数据
	_, err = client.DeleteUserSign(context.Background(), &rpc.UserAuth{Uid: newBaseData.Data["uid"]})
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("clean data ok")
	}
}

func TestMysqlCommentModel(t *testing.T) {
	InitTestENV()
	client := controller.TakeDataBaseLink()
	// 先创建用户
	userName := "xiao-hui_comment"
	userPassword := "helloworld"
	_, err := client.CreateUserSign(context.Background(), &rpc.UserSign{
		UserName:     userName,
		UserPassword: userPassword,
		SignType:     rpc.LoginAndSignType_Native,
	})
	st, _ := status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
	} else {
		t.Log("created user sign ok")
	}
	// 获得用户uid
	baseData, err := client.CheckUserStatus(context.Background(), &rpc.UserLogin{
		UserName:     userName,
		UserPassword: userPassword,
		LoginType:    rpc.LoginAndSignType_Native,
	})
	st, _ = status.FromError(err)
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
		panic(st.Message)
	} else {
		t.Log("get user id ok")
	}
	// 退出清理
	defer client.DeleteUserSign(context.Background(),&rpc.UserAuth{Uid: baseData.Data["uid"]})
	// 格式化uid
	uid, err := strconv.Atoi(baseData.Data["uid"])
	if err != nil {
		panic(err)
	}
	// 创建文章以供测试
	article, _ := client.AddArticle(context.Background(), &rpc.Article{
		ArticleAbstract: "我是摘要",
		ArticleContent:  "我是文本",
		ArticleTitle:    "我是标题",
		ArticleTag:      nil,
		Uid:             int64(uid),
	})
	// 退出后清理
	defer client.DelArticle(context.Background(),&rpc.ArticleRequest{ArticleId: []string{article.ArticleId}})
	// 添加文章主评论
	_, err = client.AddComment(context.Background(), &rpc.CommentSlave{
		Type:      rpc.CommentType_ArticleMasterComment,
		Text:      "我是第1个评论",
		Uid:       int32(uid),
		ArticleId: article.ArticleId,
	})
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
		panic(st.Message)
	} else {
		t.Log("add 1st article comment ok")
	}
	// 添加剩余十个主评论，并给每个主评论添加十个子评论
	for i := 2; i <= 11; i++ {
		_, err = client.AddComment(context.Background(), &rpc.CommentSlave{
			Type:      rpc.CommentType_ArticleMasterComment,
			Text:      fmt.Sprintf("我是第%d个评论", i),
			Uid:       int32(uid),
			ArticleId: article.ArticleId,
		})
	}
	// 添加子评论
	masterComments, err := client.GetComment(context.Background(), &rpc.CommentSlave{ArticleId: article.ArticleId})
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
		panic(st.Message)
	} else {
		t.Log(masterComments)
		t.Log("get article master comments ok")
	}
	for _, v := range masterComments.Master {
		for i := 1; i <= 10; i++ {
			_, err = client.AddComment(context.Background(), &rpc.CommentSlave{
				Type:       rpc.CommentType_ArticleSlaveComment,
				CommentMid: v.CommentMid,
				Text:       fmt.Sprintf("子评论-我是第%d个子评论", i),
				Uid:        int32(uid),
				ArticleId:  article.ArticleId,
			})
		}
	}
	// 获取添加完的评论
	comments, err := client.GetComment(context.Background(), &rpc.CommentSlave{ArticleId: article.ArticleId})
	if st.Code != 0 {
		t.Error(st.Code)
		t.Error(st.Message)
		panic(st.Message)
	} else {
		t.Log("get article comments ok")
		t.Log(comments)
	}
	// 删除主评论和子评论
	for _, v := range masterComments.Master {
		_, err := client.DeleteComment(context.Background(), &rpc.CommentSlave{
			CommentMid: v.CommentMid,
			Type:       rpc.CommentType_ArticleMasterComment,
			Text:       "删除评论",
			Uid:        v.Uid,
			ArticleId:  v.ArticleId,
		})
		st, _ := status.FromError(err)
		if st.Code != 0 {
			panic(st.Code)
		}
	}
	t.Log("del article master and slave comments step ok!")
}

package test

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/common/config"
	"com.youyu.api/common/path"
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

// 测试article数据库模型，包含子表
func TestMysqlApiArticle(t *testing.T) {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc business_server")
	}
	defer conn.Close()
	client := rpc.NewMysqlApiClient(conn)
	_, err = client.AddArticle(context.Background(), &rpc.Article{
		ArticleId:         "3",
		ArticleAbstract:   "hello worlds",
		ArticleContent:    "hello world,hello world",
		ArticleTitle:      "sb",
		ArticleTag:        nil,
		Uid:               1,
		ArticleCreateTime: time.Now().Unix(),
		ArticleUpdateTime: time.Now().Unix(),
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test addArticle ok")
	}
	_, err = client.AddArticleStatisticsHot(context.Background(), &rpc.GetArticleRequest{ArticleId: "3"})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test addArticleStatisticsHot ok")
	}

	_, err = client.AddArticleStatisticsFabulous(context.Background(), &rpc.GetArticleRequest{ArticleId: "3"})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test addArticleStatisticsFabulous ok")
	}

	_, err = client.AddArticleStatisticsCommentNum(context.Background(), &rpc.GetArticleRequest{ArticleId: "3"})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test addArticleStatisticsCommentNum ok")
	}

	article, err := client.GetArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: "3"})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(article)
		t.Log("test getArticle ok")
	}
	time.Sleep(time.Second * 2)
	rpcErr, err := client.UpdateArticle(context.Background(), &rpc.Article{
		ArticleId:         "3",
		ArticleAbstract:   "世界你好",
		ArticleContent:    "我是世界",
		ArticleTitle:      "世界是我",
		ArticleTag:        nil,
		Uid:               10,
		ArticleCreateTime: 0,
		ArticleUpdateTime: 0,
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		result, _ := client.GetArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: "3"})
		t.Log(result)
		t.Log("test updateArticle ok" + rpcErr.Message)
	}
	// 连接查询
	result, err := client.GetArticleList(context.Background(), &rpc.ArticleOptions{
		Type:    "desc",
		Order:   "hot",
		Page:    1,
		PageNum: 3,
	})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result)
		t.Log("test getArticleList ok" + rpcErr.Message)
	}
	rpcErr, err = client.DelArticle(context.Background(), &rpc.GetArticleRequest{ArticleId: "3"})
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("test delArticle ok" + rpcErr.Message)
	}
}

// 测试advertisement数据库模型
func TestMysqlApiAdvertisement(t *testing.T) {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc business_server")
	}
	defer conn.Close()
	client := rpc.NewMysqlApiClient(conn)
	testData := &rpc.Advertisement{
		AdvertisementId:     1,
		AdvertisementType:   2,
		AdvertisementLink:   "https://xiao-hui.net/NewMysqlApiClient",
		AdvertisementWeight: 9,
		AdvertisementBody:   "https://tencent/video/11",
		AdvertisementOwner:  "youyu.Inc",
	}
	_, err = client.AddAdvertisement(context.Background(), testData)
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

func TestCentApi(t *testing.T) {
	var busiess config.Config = &config.BusinessConfig{}
	r,err := busiess.GetConfig()
	if err != nil {
		t.Error("client cannot dial grpc business_server")
	}
	result := r.(*config.BusinessConfig)
	conn, err := grpc.Dial(result.DataRPCServer.IP + ":" + result.DataRPCServer.Port, grpc.WithInsecure(), grpc.WithBlock())
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
		t.Errorf("%+v",err)
	} else {
		t.Log("test SetRpcServerConfFile ok")
	}
	// business
	conf, err = client.GetBusinessConfFile(context.Background(), &rpc.Null{})
	if err != nil {
		t.Errorf("%+v",err)
	} else {
		t.Log("test SetRpcServerConfFile ok")
	}
	businessConfig := config.BusinessConfig{}
	newResult, err := businessConfig.Unmarshal(conf.Data)
	if err != nil {
		t.Errorf("%+v",err)
	} else {
		t.Log(newResult)
		t.Log("test getBusinessConfFile ok")
	}
	_, err = client.SetBusinessConfFile(context.Background(), &rpc.Config{
		Type: path.ConfBusinessRequestType,
		Data: conf.Data,
	})
	if err != nil {
		t.Errorf("%+v",err)
	} else {
		t.Log("test setBusinessConfFile ok")
	}
	client.PushLogStream(context.Background())
}
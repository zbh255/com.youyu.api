package data_rpc

import (
	"com.youyu.api/app/rpc/model"
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/lib/ecode"
	"com.youyu.api/lib/ecode/status"
	"com.youyu.api/lib/log"
	"com.youyu.api/lib/utils"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"sync"
	"time"
)

type MysqlApiServer struct {
	rpc.UnimplementedMysqlApiServer
	Logger log.Logger
	// 热度和点赞访问的互斥锁
	Lock sync.Mutex
}

// AddArticle Tag按;分割
// TODO:go-encrypt代替标准库加密的api
func (s *MysqlApiServer) AddArticle(ctx context.Context, article *rpc.Article) (*rpc.Article, error) {
	// 重试次数上限为3
	num := 0
	md := model.Article{}
	aeLoop:
	t := time.Now()
	// 文章id=md5(用户uid+文章标题+文章创建时间时间戳)
	hash := md5.New()
	hash.Write([]byte(strconv.FormatInt(article.Uid, 10) + article.ArticleTitle + strconv.FormatInt(t.Unix(), 10)))
	articleModel, err := md.AddArticle(&model.Article{
		Id:         hex.EncodeToString(hash.Sum(nil)),
		Abstract:   article.ArticleAbstract,
		Content:    article.ArticleContent,
		Title:      article.ArticleTitle,
		Tag:        utils.TagListToSplitStrings(article.ArticleTag),
		Uid:        article.Uid,
		CreateTime: t,
		UpdateTime: t,
	})
	switch errors.Cause(err) {
	case model.ArticleIdAlreadyExists:
		// article id 存在则重试
		if num <= 3 {
			num++
			goto aeLoop
		}
		return &rpc.Article{},status.Error(ecode.AddArticleErr,ecode.AddArticleErr.Message())
	case nil:
		return &rpc.Article{
			ArticleId: articleModel.Id,
		}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "add article failed"))
		return &rpc.Article{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
}

func (s *MysqlApiServer) GetArticleList(ctx context.Context, null *rpc.ArticleOptions) (*rpc.Article_Response, error) {
	md := model.Article{}
	results, err := md.GetArticles(&model.SelectOptions{
		Type:    null.Type,
		Order:   null.Order,
		Page:    null.Page,
		PageNum: null.PageNum,
	})
	if err != nil {
		s.Logger.Error(errors.Wrap(err, "get ArticleList failed"))
		return &rpc.Article_Response{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
	response := &rpc.Article_Response{Articles: make([]*rpc.ArticleLinkTab, 0)}
	for k := range results {
		response.Articles = append(response.Articles, &rpc.ArticleLinkTab{
			ArticleId:         results[k].Id,
			ArticleAbstract:   results[k].Abstract,
			ArticleTitle:      results[k].Title,
			ArticleTag:        utils.SplitStringsToTagList(results[k].Tag),
			Uid:               results[k].Uid,
			ArticleCreateTime: results[k].CreateTime.Unix(),
			ArticleUpdateTime: results[k].UpdateTime.Unix(),
			ArticleFabulous:   results[k].Fabulous,
			ArticleHot:        results[k].Hot,
			ArticleCommentNum: results[k].CommentNum,
		})
	}
	return response, nil
}

// TODO:Tag类型的问题未解决
func (s *MysqlApiServer) GetArticle(ctx context.Context, request *rpc.GetArticleRequest) (*rpc.Article, error) {
	md := model.Article{}
	article, err := md.GetArticle(request.ArticleId)

	switch errors.Cause(err) {
	case model.ArticleIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v",errors.Wrap(err,request.ArticleId)))
		return &rpc.Article{},status.Error(ecode.GetArticleErr,ecode.GetArticleErr.Message())
	case nil:
		return &rpc.Article{
			ArticleId:         article.Id,
			ArticleAbstract:   article.Abstract,
			ArticleContent:    article.Content,
			ArticleTitle:      article.Title,
			ArticleTag:        utils.SplitStringsToTagList(article.Tag),
			Uid:               article.Uid,
			ArticleCreateTime: article.CreateTime.Unix(),
			ArticleUpdateTime: article.UpdateTime.Unix(),
		}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "get article failed"))
		return &rpc.Article{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
}

func (s *MysqlApiServer) GetArticleStatistics(ctx context.Context, request *rpc.GetArticleRequest) (*rpc.ArticleStatistics, error) {
	as := model.ArticleStatistics{}
	result, err := as.GetArticleStatistics(request.ArticleId)
	switch errors.Cause(err) {
	case model.ArticleIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v",errors.Wrap(err,"article id not exist: " + request.ArticleId)))
		return &rpc.ArticleStatistics{},status.Error(ecode.GetArticleStatisticsErr,ecode.GetArticleStatisticsErr.Message())
	case nil:
		return &rpc.ArticleStatistics{
			ArticleId:         result.Id,
			ArticleFabulous:   result.Fabulous,
			ArticleHot:        result.Hot,
			ArticleCommentNum: result.CommentNum,
		}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "get article statistics failed"))
		return &rpc.ArticleStatistics{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

func (s *MysqlApiServer) AddArticleStatisticsFabulous(ctx context.Context, null *rpc.GetArticleRequest) (*rpc.Null, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	as := model.ArticleStatistics{}
	err := as.AddFabulous(null.ArticleId)
	switch errors.Cause(err) {
	case model.ArticleIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v", errors.Wrap(err, "article id not exist: "+null.ArticleId)))
		return &rpc.Null{}, status.Error(ecode.AddArticleFabulousErr, ecode.AddArticleFabulousErr.Message())
	case nil:
		return &rpc.Null{}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "add article Statistics Fabulous failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
}

func (s *MysqlApiServer) AddArticleStatisticsHot(ctx context.Context, null *rpc.GetArticleRequest) (*rpc.Null, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	as := model.ArticleStatistics{}
	err := as.AddHot(null.ArticleId)
	switch errors.Cause(err) {
	case model.ArticleIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v", errors.Wrap(err, "article id not exist: "+null.ArticleId)))
		return &rpc.Null{}, status.Error(ecode.AddArticleHotErr, ecode.AddArticleHotErr.Message())
	case nil:
		return &rpc.Null{}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "add article statistics hot failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
}

func (s *MysqlApiServer) AddArticleStatisticsCommentNum(ctx context.Context, null *rpc.GetArticleRequest) (*rpc.Null, error) {
	as := model.ArticleStatistics{}
	err := as.AddCommentNum(null.ArticleId)
	switch errors.Cause(err) {
	case model.ArticleIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v", errors.Wrap(err, "article id not exist: "+null.ArticleId)))
		return &rpc.Null{}, status.Error(ecode.AddArticleCommentNumErr, ecode.AddArticleCommentNumErr.Message())
	case nil:
		return &rpc.Null{}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "add article statistics comment num failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

func (s *MysqlApiServer) AddAdvertisement(ctx context.Context, advertisement *rpc.Advertisement) (*rpc.Null, error) {
	md := model.Advertisement{}
	err := md.AddAdvertisement(&model.Advertisement{
		Id:     advertisement.AdvertisementId,
		Type:   advertisement.AdvertisementType,
		Link:   advertisement.AdvertisementLink,
		Weight: advertisement.AdvertisementWeight,
		Body:   advertisement.AdvertisementBody,
		Owner:  advertisement.AdvertisementOwner,
	})
	if err != nil {
		s.Logger.Error(errors.Wrap(err, "add advertisement failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	} else {
		return &rpc.Null{}, nil
	}
}

func (s *MysqlApiServer) GetAdvertisement(ctx context.Context, request *rpc.AdvertisementRequest) (*rpc.Advertisement, error) {
	md := model.Advertisement{}
	result, err := md.GetAdvertisement(request.AdvertisementId)
	switch errors.Cause(err) {
	case model.AdvertisementIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v", errors.Wrap(err, "article id not exist: " + string(request.AdvertisementId))))
		return &rpc.Advertisement{}, status.Error(ecode.GetAdvertisementErr, ecode.GetAdvertisementErr.Message())
	case nil:
		return &rpc.Advertisement{
			AdvertisementId:     result.Id,
			AdvertisementType:   result.Type,
			AdvertisementLink:   result.Link,
			AdvertisementWeight: result.Weight,
			AdvertisementBody:   result.Body,
			AdvertisementOwner:  result.Owner,
		}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "get advertisement failed"))
		return &rpc.Advertisement{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

func (s *MysqlApiServer) GetAdvertisementList(ctx context.Context, null *rpc.ArticleOptions) (*rpc.AdvertisementResponse, error) {
	md := model.Advertisement{}
	results, err := md.GetAdvertisements(&model.SelectOptions{
		Type:    null.Type,
		Order:   null.Order,
		Page:    null.Page,
		PageNum: null.PageNum,
	})
	if err != nil {
		s.Logger.Error(errors.Wrap(err, "get advertisement list failed"))
		return &rpc.AdvertisementResponse{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
	response := &rpc.AdvertisementResponse{AdvertisementList: make([]*rpc.Advertisement, 0)}
	for k := range results {
		response.AdvertisementList = append(response.AdvertisementList, &rpc.Advertisement{
			AdvertisementId:     results[k].Id,
			AdvertisementType:   results[k].Type,
			AdvertisementLink:   results[k].Link,
			AdvertisementWeight: results[k].Weight,
			AdvertisementBody:   results[k].Body,
			AdvertisementOwner:  results[k].Owner,
		})
	}
	return response, nil
}

func (s *MysqlApiServer) UpdateArticle(ctx context.Context, article *rpc.Article) (*rpc.Null, error) {
	md := model.Article{}
	err := md.SetArticle(&model.Article{
		Id:         article.ArticleId,
		Abstract:   article.ArticleAbstract,
		Content:    article.ArticleContent,
		Title:      article.ArticleTitle,
		Tag:        utils.TagListToSplitStrings(article.ArticleTag),
		Uid:        article.Uid,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	switch errors.Cause(err) {
	case model.ArticleIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v", errors.Wrap(err, "article id not exist: " + string(article.ArticleId))))
		return &rpc.Null{}, status.Error(ecode.UpdArticleErr, ecode.UpdArticleErr.Message())
	case nil:
		return &rpc.Null{}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "update article failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
}

func (s *MysqlApiServer) DelArticle(ctx context.Context, request *rpc.GetArticleRequest) (*rpc.Null, error) {
	md := model.Article{}
	err := md.DelArticle(request.ArticleId)
	if err != nil {
		s.Logger.Error(errors.Wrap(err, "del article failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	} else {
		return &rpc.Null{}, nil
	}
}

func (s *MysqlApiServer) DelArticleStatisticsFabulous(ctx context.Context, request *rpc.GetArticleRequest) (*rpc.Null, error) {
	md := model.ArticleStatistics{}
	err := md.ReduceFabulous(request.ArticleId)
	switch errors.Cause(err) {
	case model.ArticleIdNotExists:
		s.Logger.Info(fmt.Sprintf("%+v", errors.Wrap(err, "article id not exist: "+request.ArticleId)))
		return &rpc.Null{}, status.Error(ecode.DelArticleFabulousErr, ecode.DelArticleFabulousErr.Message())
	case nil:
		return &rpc.Null{}, nil
	default:
		s.Logger.Error(errors.Wrap(err, "del article statistics fabulous failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

func (s *MysqlApiServer) UpdateAdvertisement(ctx context.Context, request *rpc.Advertisement) (*rpc.Null, error) {
	md := model.Advertisement{}
	err := md.SetAdvertisement(&model.Advertisement{
		Id:     request.AdvertisementId,
		Type:   request.AdvertisementType,
		Link:   request.AdvertisementLink,
		Weight: request.AdvertisementWeight,
		Body:   request.AdvertisementBody,
		Owner:  request.AdvertisementOwner,
	})
	if err != nil {
		s.Logger.Error(errors.Wrap(err, "update advertisement failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	} else {
		return &rpc.Null{}, nil
	}
}

func (s *MysqlApiServer) DelAdvertisement(ctx context.Context, request *rpc.AdvertisementRequest) (*rpc.Null, error) {
	md := model.Advertisement{}
	err := md.DelAdvertisement(request.AdvertisementId)
	if err != nil {
		s.Logger.Error(errors.Wrap(err, "del advertisement failed"))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	} else {
		return &rpc.Null{}, nil
	}
}

// 根据id获取Tag文本 支持批量获取
func (s *MysqlApiServer) GetTagText(ctx context.Context, tag *rpc.Tag) (*rpc.Tag, error) {
	md := model.Tags{}
	text, err := md.GetTagText(tag.Tid)
	switch errors.Cause(err) {
	case model.TagIdNotExists:
		return &rpc.Tag{},status.Error(ecode.TagIdNotExists,ecode.TagIdNotExists.Message())
	case nil:
		tidList := make([]int32,len(text))
		textList := make([]string,len(text))
		for k := range text {
			tidList[k] = text[k].Tid
			textList[k] = text[k].Text
		}
		return &rpc.Tag{Tid: tidList, Text: textList}, nil
	default:
		s.Logger.Error(err)
		return &rpc.Tag{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

// 添加Tag 一次只能添加一个
func (s *MysqlApiServer) AddTag(ctx context.Context, tag *rpc.Tag) (*rpc.Null, error) {
	md := model.Tags{}
	if tag.Text == nil || len(tag.Text) == 0{
		return &rpc.Null{},status.Error(ecode.TagNameAlreadyExists,ecode.TagNameAlreadyExists.Message())
	}
	err := md.AddTag(tag.Text[0])
	switch errors.Cause(err) {
	case model.TagNameAlreadyExists:
		return &rpc.Null{},status.Error(ecode.TagNameAlreadyExists,ecode.TagNameAlreadyExists.Message())
	case nil:
		return &rpc.Null{}, nil
	default:
		s.Logger.Error(err)
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

// NOTE: 非给普通用户客户端开放的接口
// 根据Id 删除Tag,不支持批量操作
func (s *MysqlApiServer) DelTag(ctx context.Context, tag *rpc.Tag) (*rpc.Null, error) {
	md := model.Tags{}
	if tag.Tid == nil || len(tag.Tid) == 0 {
		return &rpc.Null{}, status.Error(ecode.TagIdNotExists,ecode.TagIdNotExists.Message())
	}
	err := md.DelTag(tag.Tid[0])
	if err != nil {
		s.Logger.Error(err)
		return &rpc.Null{}, err
	} else {
		return &rpc.Null{}, nil
	}
}

// 根据文本获取Tag Id支持批量获取
func (s *MysqlApiServer) GetTagInt32Id(ctx context.Context, tag *rpc.Tag) (*rpc.Tag, error) {
	md := model.Tags{}
	int32Id, err := md.GetTagInt32Id(tag.Text)
	switch errors.Cause(err) {
	case model.TagNameNotExists:
		return &rpc.Tag{},status.Error(ecode.TagNameNotExists,ecode.TagNameNotExists.Message())
	case nil:
		tidList := make([]int32,len(int32Id))
		textList := make([]string,len(int32Id))
		for k := range int32Id {
			tidList[k] = int32Id[k].Tid
			textList[k] = int32Id[k].Text
		}
		return &rpc.Tag{Tid: tidList, Text: textList}, nil
	default:
		s.Logger.Error(err)
		return &rpc.Tag{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

// TODO: 接入微信的注册登录接口
// TODO: 接入第三方平台的滑动验证
// 创建一个用户
func (s *MysqlApiServer) CreateUserSign(ctx context.Context, sign *rpc.UserLoginOrSign) (*rpc.Null, error) {
	md := model.UserBase{}
	// 创建用户账户信息
	// 创建用户信息
	// level 9为普通用户
	_, err := md.CreateUser(&model.UserBase{
		UserPassword: sign.UserPassword,
		Name:         sign.UserName,
	})
	switch errors.Cause(err) {
	case model.UserNameAlreadyExists:
		s.Logger.Error(err)
		return &rpc.Null{}, status.Error(ecode.UserDuplicate,ecode.UserDuplicate.Message())
	case nil:
		return &rpc.Null{}, nil
	default:
		s.Logger.Error(err)
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
}

// 获取用户信息
func (s *MysqlApiServer) GetUserInfo(ctx context.Context, auth *rpc.UserAuth) (*rpc.UserInfoShow, error) {
	md := model.UserInfo{}
	uid, err := strconv.Atoi(auth.Uid)
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.UserInfoShow{}, err
	}
	userInfo, err := md.GetUserInfo(int32(uid))
	switch errors.Cause(err) {
	case model.UserDoesNotExist:
		return &rpc.UserInfoShow{},status.Error(ecode.UserNotExist,ecode.UserNotExist.Message())
	case nil:
		return &rpc.UserInfoShow{
			Uid:          userInfo.Uid,
			Level:        userInfo.Level,
			Phone:        userInfo.Phone,
			Email:        userInfo.Email,
			PhoneStatus:  int32(userInfo.PhoneStatus),
			EmailStatus:  int32(userInfo.EmailStatus),
			CreateTime:   userInfo.CreateTime.String(),
			Sex:          int32(userInfo.Sex),
			Age:          int32(userInfo.Age),
			UserName:     userInfo.Name,
			UserNickName: userInfo.NickName,
		}, nil
	default:
		return &rpc.UserInfoShow{},status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}

}

// 更新用户信息
func (s *MysqlApiServer) UpdateUserInfo(ctx context.Context, set *rpc.UserInfoSet) (*rpc.Null, error) {
	md := model.UserInfo{}
	err := md.UpdateUserInfo(&model.UserInfo{
		Uid:      set.Uid,
		Phone:    set.Phone,
		Email:    set.Email,
		Sex:      int(set.Sex),
		Age:      int(set.Age),
		NickName: set.UserNickName,
		Addr:     set.Addr,
		Explain:  set.Explain,
	})
	if err != nil {
		s.Logger.Error(err)
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
	return &rpc.Null{}, nil
}

// 删除用户
func (s *MysqlApiServer) DeleteUserSign(ctx context.Context, sign *rpc.UserAuth) (*rpc.Null, error) {
	md := model.UserBase{}
	uid, err := strconv.Atoi(sign.Uid)
	if err != nil {
		s.Logger.Error(errors.WithStack(err))
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
	err = md.DeleteUser(int32(uid))
	if err != nil {
		s.Logger.Error(err)
		return &rpc.Null{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
	return &rpc.Null{}, nil
}

// 验证成功之后会返回uid,user_name供调用者签钥
// err格式 code : 0 -> message : 具体的业务错误码
func (s *MysqlApiServer) CheckUserStatus(ctx context.Context, sign *rpc.UserLoginOrSign) (*rpc.BaseData, error) {
	md := model.UserBase{}
	err := md.CheckUser(sign.UserName, sign.UserPassword)
	switch errors.Cause(err) {
	case model.UserDoesNotExist:
		s.Logger.Error(err)
		return &rpc.BaseData{}, status.Error(ecode.UserNotExist, ecode.UserNotExist.Message())
	case model.UserPasswordORUserNameErr:
		return &rpc.BaseData{}, status.Error(ecode.UsernameOrPasswordErr,ecode.UsernameOrPasswordErr.Message())
	case nil:
		return &rpc.BaseData{
			Data: map[string]string{
				"uid":       strconv.Itoa(int(md.Uid)),
				"user_name": md.Name,
			},
		}, nil
	default:
		return &rpc.BaseData{}, status.Error(ecode.ServerErr,ecode.ServerErr.Message())
	}
}

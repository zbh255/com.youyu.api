package server

import (
	"com.youyu.api/app/rpc/model"
	rpc "com.youyu.api/app/rpc/proto_files"
	"com.youyu.api/common/errors"
	"com.youyu.api/common/utils"
	"context"
	"time"
)

type MysqlApiServer struct {
	rpc.UnimplementedMysqlApiServer
}

func (s *MysqlApiServer) AddArticle(ctx context.Context, article *rpc.Article) (*rpc.Errors, error) {
	md := model.Article{}
	t := time.Now()

	err := md.AddArticle(&model.Article{
		Id:         article.ArticleId,
		Abstract:   article.ArticleAbstract,
		Content:    article.ArticleContent,
		Title:      article.ArticleTitle,
		Tag:        "",
		Uid:        article.Uid,
		CreateTime: t,
		UpdateTime: t,
	})
	if err != nil {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, &errors.Errno{
			Code:     500,
			Message:  err.Error(),
			HttpCode: 500,
		}), err
	}
	return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
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
		return nil, err
	}
	response := &rpc.Article_Response{Articles: make([]*rpc.ArticleLinkTab, 0)}
	for k := range results {
		response.Articles = append(response.Articles, &rpc.ArticleLinkTab{
			ArticleId:         results[k].Id,
			ArticleAbstract:   results[k].Abstract,
			ArticleTitle:      results[k].Title,
			ArticleTag:        nil,
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
	if err != nil {
		return nil, err
	}
	return &rpc.Article{
		ArticleId:         article.Id,
		ArticleAbstract:   article.Abstract,
		ArticleContent:    article.Content,
		ArticleTitle:      article.Title,
		ArticleTag:        nil,
		Uid:               article.Uid,
		ArticleCreateTime: article.CreateTime.Unix(),
		ArticleUpdateTime: article.UpdateTime.Unix(),
	}, nil
}

func (s *MysqlApiServer) GetArticleStatistics(ctx context.Context, request *rpc.GetArticleRequest) (*rpc.ArticleStatistics, error) {
	as := model.ArticleStatistics{}
	result, err := as.GetArticleStatistics(request.ArticleId)
	if err != nil {
		return nil, err
	} else {
		return &rpc.ArticleStatistics{
			ArticleId:         result.Id,
			ArticleFabulous:   result.Fabulous,
			ArticleHot:        result.Hot,
			ArticleCommentNum: result.CommentNum,
		}, err
	}
}

func (s *MysqlApiServer) AddArticleStatisticsFabulous(ctx context.Context, null *rpc.GetArticleRequest) (*rpc.Errors, error) {
	as := model.ArticleStatistics{}
	err := as.AddFabulous(null.ArticleId)
	if err != nil {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, &errors.Errno{
			Code:     500,
			Message:  err.Error(),
			HttpCode: 500,
		}), err
	}
	return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
}

func (s *MysqlApiServer) AddArticleStatisticsHot(ctx context.Context, null *rpc.GetArticleRequest) (*rpc.Errors, error) {
	as := model.ArticleStatistics{}
	err := as.AddHot(null.ArticleId)
	if err != nil {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, &errors.Errno{
			Code:     500,
			Message:  err.Error(),
			HttpCode: 500,
		}), err
	}
	return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
}

func (s *MysqlApiServer) AddArticleStatisticsCommentNum(ctx context.Context, null *rpc.GetArticleRequest) (*rpc.Errors, error) {
	as := model.ArticleStatistics{}
	err := as.AddCommentNum(null.ArticleId)
	if err != nil {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, &errors.Errno{
			Code:     500,
			Message:  err.Error(),
			HttpCode: 500,
		}), err
	}
	return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
}

func (s *MysqlApiServer) AddAdvertisement(ctx context.Context, advertisement *rpc.Advertisement) (*rpc.Errors, error) {
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
		return nil, err
	} else {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
	}
}

func (s *MysqlApiServer) GetAdvertisement(ctx context.Context, request *rpc.AdvertisementRequest) (*rpc.Advertisement, error) {
	md := model.Advertisement{}
	result, err := md.GetAdvertisement(request.AdvertisementId)
	if err != nil {
		return nil, err
	} else {
		return &rpc.Advertisement{
			AdvertisementId:     result.Id,
			AdvertisementType:   result.Type,
			AdvertisementLink:   result.Link,
			AdvertisementWeight: result.Weight,
			AdvertisementBody:   result.Body,
			AdvertisementOwner:  result.Owner,
		}, nil
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
		return nil, err
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

func (s *MysqlApiServer) UpdateArticle(ctx context.Context, article *rpc.Article) (*rpc.Errors, error) {
	md := model.Article{}
	err := md.SetArticle(&model.Article{
		Id:         article.ArticleId,
		Abstract:   article.ArticleAbstract,
		Content:    article.ArticleContent,
		Title:      article.ArticleTitle,
		Tag:        "",
		Uid:        article.Uid,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	} else {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
	}
}

func (s *MysqlApiServer) DelArticle(ctx context.Context, request *rpc.GetArticleRequest) (*rpc.Errors, error) {
	md := model.Article{}
	err := md.DelArticle(request.ArticleId)
	if err != nil {
		return nil, err
	} else {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
	}
}

func (s *MysqlApiServer) DelArticleStatisticsFabulous(ctx context.Context, request *rpc.GetArticleRequest) (*rpc.Errors, error) {
	md := model.ArticleStatistics{}
	err := md.ReduceFabulous(request.ArticleId)
	if err != nil {
		return nil, err
	} else {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
	}
}

func (s *MysqlApiServer) UpdateAdvertisement(ctx context.Context, request *rpc.Advertisement) (*rpc.Errors, error) {
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
		return nil, err
	} else {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
	}
}

func (s *MysqlApiServer) DelAdvertisement(ctx context.Context, request *rpc.AdvertisementRequest) (*rpc.Errors, error) {
	md := model.Advertisement{}
	err := md.DelAdvertisement(request.AdvertisementId)
	if err != nil {
		return nil, err
	} else {
		return utils.CustomErrToGrpcCustomErr(&rpc.Errors{}, errors.OK), nil
	}
}

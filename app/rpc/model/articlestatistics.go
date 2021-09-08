// Package model Table Articlestatistics绑定的操作方法
package model

import "github.com/pkg/errors"

func (self *ArticleStatistics) AddFabulous(id string) error {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_fabulous", self.Fabulous+1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// 点赞-1
func (self *ArticleStatistics) ReduceFabulous(id string) error {
	//db, err := DB.GetConnect()
	//if err != nil {
	//	return errors.Wrap(err, "")
	//}
	//if db == nil {
	//	return errors.New("mysql server connection failed")
	//}
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_fabulous", self.Fabulous-1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

// 获得存储文章点赞等数据的表
func (self *ArticleStatistics) GetArticleStatistics(id string) (*ArticleStatistics, error) {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return nil, errors.New("the query record is zero")
	} else {
		return self, nil
	}
}

// 错误类型为数据库连接错误和查询错误
func (self *ArticleStatistics) AddHot(id string) error {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_hot", self.Hot+1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (self *ArticleStatistics) AddCommentNum(id string) error {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.New("the query record is zero")
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_comment_num", self.CommentNum+1)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (self *ArticleStatistics) AddSelf(statistics *ArticleStatistics) error {
	result := DB.Create(statistics)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

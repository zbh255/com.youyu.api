// Package model Table Articlestatistics绑定的操作方法
package model

import "github.com/pkg/errors"

// AddFabulous 添加文章点赞数
// pkg/errors处理错误
// article_id 不存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
func (self *ArticleStatistics) AddFabulous(id string) error {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.WithStack(ArticleIdNotExists)
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_fabulous", self.Fabulous+1)
	if result.Error != nil {
		return errors.WithStack(result.Error)
	} else {
		return nil
	}
}

// ReduceFabulous 点赞-1
// pkg/errors处理错误
// article_id 不存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
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
		return errors.WithStack(ArticleIdNotExists)
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_fabulous", self.Fabulous-1)
	if result.Error != nil {
		return errors.WithStack(result.Error)
	} else {
		return nil
	}
}

// 获得存储文章点赞等数据的表
func (self *ArticleStatistics) GetArticleStatistics(id string) (*ArticleStatistics, error) {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(ArticleIdNotExists)
	} else {
		return self, result.Error
	}
}

// 错误类型为数据库连接错误和查询错误
// pkg/errors处理错误
// article_id 不存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
func (self *ArticleStatistics) AddHot(id string) error {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.WithStack(ArticleIdNotExists)
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_hot", self.Hot+1)
	if result.Error != nil {
		return errors.WithStack(result.Error)
	} else {
		return nil
	}
}

// pkg/errors处理错误
// article_id 不存在的时候返回自包装错误
// 操作数据时有错误则返回gorm 的原始错误
func (self *ArticleStatistics) AddCommentNum(id string) error {
	result := DB.Where("article_id = ?", id).First(self)
	// 判断查找到的记录
	if result.RowsAffected == 0 {
		return errors.WithStack(ArticleIdNotExists)
	}
	result = DB.Model(self).Where("article_id = ?", id).Update("article_comment_num", self.CommentNum+1)
	if result.Error != nil {
		return errors.WithStack(result.Error)
	} else {
		return nil
	}
}

func (self *ArticleStatistics) AddSelf(statistics *ArticleStatistics) error {
	result := DB.Create(statistics)
	if result.Error != nil {
		return errors.WithStack(result.Error)
	} else {
		return nil
	}
}

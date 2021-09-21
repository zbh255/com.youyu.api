package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (t *Tags) AddTag(text string) error {
	//db, err := DB.GetConnect()
	//if err != nil {
	//	return errors.WithStack(err)
	//}
	//if db == nil {
	//	return errors.New("mysql server connection failed")
	//}
	t.Text = text
	// 标签的文字不能相同
	result := DB.Where("text = ?", t.Text).First(t)
	if result.RowsAffected > 0 {
		return errors.WithStack(TagNameAlreadyExists)
	} else {
		return errors.WithStack(DB.Create(t).Error)
	}
	//if num := DB.Where("text = ?", text).First(t).RowsAffected; num > 0 {
	//	return errors.Wrap(DB.Create(t).Error, "add tag failed")
	//} else {
	//	return errors.Wrap(CreateSameExistence, "tag text cannot be the same")
	//}
}

// GetTagText pkg/errors处理错误
// tid 不存在时返回自定义错误
// 操作数据时有错误则返回gorm 的原始错误
func (t *Tags) GetTagText(tid []int32) (tags []*Tags, err error) {
	if tid == nil || len(tid) == 0 {
		return nil, errors.WithStack(TagIdNotExists)
	}
	result := DB.Where("tid IN ?",tid).Find(&tags)
	if result.RowsAffected == 0 {
		return nil,errors.WithStack(TagIdNotExists)
	} else {
		return tags, errors.WithStack(result.Error)
	}
}

// GetTagInt32Id pkg/errors处理错误
// text 不存在时返回自定义错误
// 操作数据时有错误则返回gorm 的原始错误
func (t *Tags) GetTagInt32Id(text []string) (tags []*Tags, err error) {
	if text == nil || len(text) == 0 {
		return nil, errors.WithStack(TagNameNotExists)
	}
	result := DB.Where("text IN ?", text).Find(&tags)
	if result.RowsAffected == 0 {
		return nil, errors.WithStack(TagNameNotExists)
	} else {
		return tags, errors.WithStack(result.Error)
	}

}

// NOTE: 非给用户客户端开放的接口，慎用
func (t *Tags) DelTag(tid int32) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		// 根据主键删除数据
		// 因为设置的外键，所以存储点赞等信息的也会自动删除
		t.Tid = tid
		if err := tx.Delete(t).Error; err != nil {
			return err
		}
		return nil
	})
	return errors.Wrap(err, "del tag failed")
}

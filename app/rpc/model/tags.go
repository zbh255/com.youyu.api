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
	if num := DB.Where("text = ?", text).First(t).RowsAffected; num == 0 {
		return errors.Wrap(DB.Create(t).Error, "add tag failed")
	} else {
		return errors.Wrap(CreateSameExistence, "tag text cannot be the same")
	}
}

func (t *Tags) GetTagText(tid int32) (string, error) {
	return t.Text, errors.Wrap(DB.Where("tid = ?", tid).First(t).Error, "get tag text failed")
}

func (t *Tags) GetTagInt32Id(text string) (int32, error) {
	//db, err := DB.GetConnect()
	//if err != nil {
	//	return -1, errors.WithStack(err)
	//}
	//if db == nil {
	//	return -1, errors.New("mysql server connection failed")
	//}
	return t.Tid, errors.Wrap(DB.Where("text = ?", text).First(t).Error, "get tag int32 id failed")
}

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

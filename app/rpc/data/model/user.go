package model

import (
	encrypt "github.com/abingzo/go-encrypt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

/*
CreateUser 注意:本方法会在数据库创建所有有关与User的信息
CreateUser 在同一个事务内,通过UserInfo模型操作数据库
CreateUser pkg/errors处理错误
*/
func (ub *UserBase) CreateUser(userBase *UserBase, info *UserInfo) (*UserBase, error) {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("name = ?", userBase.Name).First(ub).RowsAffected > 0 {
			return errors.WithStack(UserNameAlreadyExists)
		}
		userBase.Uid = 0
		t := time.Now()
		rand.Seed(t.UnixNano())
		// 为用户生成盐
		// 生成盐的方式 2^64范围之内的随机数+时间字符串+用户名
		saltBytes, err := encrypt.NewCoder().GetAbstract().Md5Coder(encrypt.HEX).
			Append(t.String()).Append(strconv.FormatInt(rand.Int63(), 10)).SumString(userBase.Name).Result()
		if err != nil {
			return errors.WithStack(err)
		}
		// 结合盐值为用户的密码生成摘要
		password, err := encrypt.NewCoder().GetAbstract().Sha512Coder(encrypt.HEX).Append(string(saltBytes)).
			SumString(userBase.UserPassword).Result()
		if err != nil {
			return errors.WithStack(err)
		}
		// 赋值并存入数据库
		userBase.Salt = string(saltBytes)
		userBase.UserPassword = string(password)
		err = tx.Create(userBase).Error
		if err != nil {
			return errors.WithStack(err)
		}
		err = tx.Where("name = ?", userBase.Name).First(ub).Error
		if err != nil {
			return errors.WithStack(err)
		}
		// info为空时创建默认用户信息表
		// 否则只修改部分数据
		if info == nil {
			userInfo := DefaultUserInfoTemplate
			userInfo.Uid = ub.Uid
			userInfo.Name = ub.Name
			userInfo.CreateTime = t
			userInfo.UpdateTime = t
			info = &userInfo
		} else {
			// TODO:uid和user_name千万不要忘了修改
			info.Uid = ub.Uid
			info.Name = ub.Name
			info.CreateTime = t
			info.UpdateTime = t
		}
		// 在同一个事务内写入数据库
		return errors.WithStack(info.CreateUserInfo(*info, tx))
	})
	return ub, err
}

// CheckUser 验证用户名和密码是否正确
// password为原生密码
func (ub *UserBase) CheckUser(userName string, userPassword string) error {
	// 获得用户的盐值
	if DB.Where("name = ?", userName).First(&ub).RowsAffected == 0 {
		return errors.WithStack(UserDoesNotExist)
	}
	// 算出用户提交的密码摘要
	result, err := encrypt.NewCoder().GetAbstract().Sha512Coder(encrypt.HEX).
		SumString(userPassword).Append(ub.Salt).Result()
	if err != nil {
		return errors.WithStack(err)
	}
	if DB.Where("name = ? AND password = ?", userName, string(result)).First(ub).RowsAffected == 0 {
		return errors.WithStack(UserPasswordORUserNameErr)
	}
	return nil
}

// 删除用户
func (ub *UserBase) DeleteUser(uid int32) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 因为设置的外键，所以的用户信息的也会自动删除
		ub.Uid = uid
		return errors.WithStack(tx.Delete(ub).Error)
	})
}

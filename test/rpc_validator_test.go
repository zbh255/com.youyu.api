// 测试grpc的参数验证器
package test

import (
	"com.youyu.api/app/rpc/proto_files"
	"testing"
)

func TestUserInfo(t *testing.T) {
	uis := proto_files.UserInfoSet{}
	err := uis.Validate()
	if err != nil {
		t.Error(err)
	}
	uis2 := proto_files.UserInfoSet{
		Uid:          0,
		Sex:          1,
		Age:          100,
		UserNickName: "wuyin",
		Explain:      "hello world",
		Country:      "中国",
		Province:     "广东",
		City:         "清远",
		DetailAddr:   "清远区清远镇清远村1栋1号",
		Language:     "zh_CN",
	}
	err = uis2.Validate()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("check ok")
	}
}
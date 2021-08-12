package test

import (
	"com.youyu.api/common/config"
	"testing"
)

func TestConfig(t *testing.T) {
	var business config.Config = &config.BusinessConfig{}
	businessConf,err := business.GetConfig()
	// TODO:优化错误处理
	result := businessConf.(*config.BusinessConfig)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log(result)
	}
	// 序列化和反序列化
	byteData := result.Marshal()
	if err != nil {
		t.Errorf("%+v",err)
	} else {
		t.Log(string(byteData))
	}
	newResult, err := result.Unmarshal(byteData)
	if err != nil {
		t.Errorf("%+v",err)
	} else {
		t.Log(newResult.(*config.BusinessConfig))
	}
}

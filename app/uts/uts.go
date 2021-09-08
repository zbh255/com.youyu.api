// Package uts 该包负责管理项目的会用到的工具
package uts

import (
	"bytes"
	"com.youyu.api/lib/alg"
	"com.youyu.api/lib/ecode"
	p "com.youyu.api/lib/path"
	"com.youyu.api/lib/utils/version"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
)

// GetAppInfoToJson 获得Uts编译时注入的一些关于应用程序的信息
func GetAppInfoToJson() {
	appInfo := version.GetInfo()
	result, err := json.Marshal(appInfo)
	if err != nil {
		log.Fatal(err)
	}
	out := bytes.Buffer{}
	err = json.Indent(&out, result, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	_, err = out.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateBusinessErrInfoToJson 创建Json格式的业务错误代码对应的信息
// path 未指定则采用lib/path下的默认值
func CreateBusinessErrInfoToJson(path string) {
	if path == "" {
		path = p.ErrMsgJsonFileName
	}
	codes := ecode.GetRegisCodes()
	// 按照绝对值大小排序
	absLessZero := make([]int,0)
	absMoreZero := make([]int,0)
	for _,v := range codes {
		if v < 0 {
			absLessZero = append(absLessZero,int(math.Abs(float64(v))))
		} else {
			absMoreZero = append(absMoreZero,v)
		}
	}
	alg.QuickSort(absLessZero)
	alg.QuickSort(absMoreZero)
	// 将绝对值复原
	for k := range absLessZero {
		absLessZero[k] = absLessZero[k] - absLessZero[k] * 2
	}
	// 拼接
	absLessZero = append(absLessZero,absMoreZero...)
	raw := make(map[int]string)
	for _, v := range absLessZero {
		raw[v] = ""
	}
	// 写入文件
	indent, err := json.MarshalIndent(raw, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(ioutil.WriteFile(p.InfoFileDefaultPath+"/"+path, indent, 0755))
}

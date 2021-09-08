package utils

import (
	rpc "com.youyu.api/app/rpc/proto_files"
	err "com.youyu.api/lib/errors"
	"com.youyu.api/lib/path"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CustomErrToGrpcCustomErr 从go文件中定义的errors转换为grpc中自定义的errors
func CustomErrToGrpcCustomErr(grpcErr *rpc.Errors, errno *err.Errno) *rpc.Errors {
	grpcErr.Code = int32(errno.Code)
	grpcErr.Message = errno.Message
	grpcErr.HttpCode = int32(errno.HttpCode)
	return grpcErr
}

func GrpcCustomErrToCustomErr(grpcErr *rpc.Errors, errno *err.Errno) *err.Errno {
	errno.Code = int(grpcErr.Code)
	errno.HttpCode = int(grpcErr.HttpCode)
	errno.Message = grpcErr.Message
	return errno
}

// ReadErrJsonToCodesMap 从io.Reader接口读取并返回注册错误信息用的Codes
// path为空则使用lib/path下的默认值
func ReadErrJsonToCodesMap(reader io.Reader) (map[int]string, error) {
	file, err2 := ioutil.ReadAll(reader)
	if err2 != nil {
		return nil, err2
	}
	stru := make(map[int]string)
	return stru, json.Unmarshal(file, &stru)
}

// WriteErrCodesMapToFile 将data中的数据写入Err_msg.json中
// path为空则使用默认path
func WriteErrCodesMapToFile(data []byte, ph string) error {
	p := path.InfoFileDefaultPath + "/" + path.ErrMsgJsonFileName
	if ph != "" {
		p = ph
	}
	return ioutil.WriteFile(p, data, 0755)
}

// TagListToSplitStrings 从一个string切片或数组格式化为使用;分割的字符串
func TagListToSplitStrings(tags []string) string {
	if tags == nil {
		return ""
	}
	return strings.Join(tags, ";")
}

// SplitStringsToTagList 从一个使用;分割的tag string转换为string切片
func SplitStringsToTagList(tag string) []string {
	if tag == "" {
		return nil
	}
	return strings.Split(tag, ";")
}

// 获取一个随机的jwt签名密钥
// 返回的string为hex(16进制)字符串
func CreateSigningKey(uid string) string {
	hash := sha256.New()
	// 时间戳加随机数生成
	t := strconv.FormatInt(time.Now().UnixNano(), 10)
	rand.Seed(time.Now().UnixNano())
	hash.Write([]byte(uid + t + strconv.FormatInt(rand.Int63(), 10)))
	return hex.EncodeToString(hash.Sum(nil))
}

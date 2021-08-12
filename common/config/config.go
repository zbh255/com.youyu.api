package config

import (
	"com.youyu.api/common/path"
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"io"
	"os"
)

// 错误定义
// 解码失败
var ConfDecodeErr = errors.New("toml decoding failed")
// 编码失败
var ConfEncodeErr = errors.New("toml encoding failed")
// 序列化
var ConfUnmarshalErr = errors.New("toml unmarshal failed")
// 反序列化
var ConfMarshalErr = errors.New("toml marshal failed")

type AutoGenerated struct {
	Database struct {
		IPAddr       string `toml:"ip_addr"`
		Port         string `toml:"port"`
		UserName     string `toml:"user_name"`
		UserPassword string `toml:"user_password"`
	} `toml:"database"`
	Server struct {
		IPAddr   string `toml:"ip_addr"`
		Port     string `toml:"port"`
		Protocol string `toml:"protocol"`
	} `toml:"server"`
}

type BusinessConfig struct {
	DataRPCServer struct {
		IP   string `toml:"ip"`
		Port string `toml:"port"`
	} `toml:"data_rpc_server"`
	CentRPCServer struct {
		IP   string `toml:"ip"`
		Port string `toml:"port"`
	} `toml:"cent_rpc_server"`
}

type IOWrite struct {
	Data []byte
}

func (i *IOWrite)Write(p []byte) (int, error) {
	i.Data = p
	return 0,nil
}

type Config interface {
	GetConfig() (interface{},error)
	SetConfig(interface{}) error
	Unmarshal([]byte) (interface{}, error)
	Marshal() []byte
}

func (ag *AutoGenerated)GetConfig() (interface{},error) {
	localPath,_ := os.Getwd()
	_, err := toml.DecodeFile(localPath+"/"+path.ConfFilePath+"/"+path.ConfRpcServerFileName, ag)
	if err != nil {
		return nil,errors.Wrap(err,ConfDecodeErr.Error())
	}
	return ag, err
}

func (ag *AutoGenerated)SetConfig(conf interface{}) error {
	localPath,_ := os.Getwd()
	result := conf.(*AutoGenerated)
	file, err := os.Create(localPath + "/" + path.ConfFilePath + "/" + path.ConfRpcServerFileName)
	if err != nil {
		return err
	}
	err = toml.NewEncoder(file).Encode(result)
	if err != nil {
		return errors.Wrap(err,ConfEncodeErr.Error())
	}
	return nil
}

// 返回github.com/pkg/errors包的自定义错误
func (ag *AutoGenerated) Unmarshal(bytes []byte) (interface{}, error) {
	err := toml.Unmarshal(bytes, ag)
	if err != nil {
		return nil, errors.Wrap(err,ConfUnmarshalErr.Error())
	}
	return ag, nil
}

// i为结构体的指针类型
// 返回值为[]bye和错误
// 错误返回项保留
func (ag *AutoGenerated) Marshal() []byte {
	struck := IOWrite{}
	var ioW io.Writer = &struck
	_ = toml.NewEncoder(ioW).Encode(ag)
	return struck.Data
}

func (b *BusinessConfig)GetConfig() (interface{},error) {
	localPath,_ := os.Getwd()
	_, err := toml.DecodeFile(localPath+"/"+path.ConfFilePath+"/"+path.ConfBusinessFileName, b)
	if err != nil {
		return nil,errors.Wrap(err,ConfDecodeErr.Error())
	}
	return b, err
}

func (b *BusinessConfig)SetConfig(conf interface{}) error {
	localPath,_ := os.Getwd()
	result := conf.(*BusinessConfig)
	file, err := os.Create(localPath + "/" + path.ConfFilePath + "/" + path.ConfBusinessFileName)
	if err != nil {
		return err
	}
	err = toml.NewEncoder(file).Encode(result)
	if err != nil {
		return errors.Wrap(err,ConfEncodeErr.Error())
	}
	return nil
}

func (b *BusinessConfig) Unmarshal(bytes []byte) (interface{}, error) {
	err := toml.Unmarshal(bytes, b)
	if err != nil {
		return nil, errors.Wrap(err,ConfUnmarshalErr.Error())
	}
	return b, nil
}

// business为结构体的指针类型
// 返回值为[]bye
// 错误返回项保留
func (b *BusinessConfig) Marshal() []byte {
	struck := IOWrite{}
	var ioW io.Writer = &struck
	_ = toml.NewEncoder(ioW).Encode(b)
	return struck.Data
}
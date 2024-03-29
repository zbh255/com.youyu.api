package config

import (
	"com.youyu.api/lib/path"
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
	DataRPCServer struct {
		IP   string `toml:"ip"`
		Port string `toml:"port"`
	} `toml:"data_rpc_server"`
	SecretKeyRPCServer struct {
		IP   string `toml:"ip"`
		Port string `toml:"port"`
	} `toml:"secret_key_rpc_server"`
	Database struct {
		IPAddr       string `toml:"ip_addr"`
		Port         string `toml:"port"`
		UserName     string `toml:"user_name"`
		UserPassword string `toml:"user_password"`
		DatabaseName string `toml:"database_name"`
		Sync         struct {
			DbMaxIdleSize     int `toml:"db_max_idle_size"`
			DbMaxIdleLifeTime int `toml:"db_max_idle_life_time"`
			DbMaxOpenConnSize int `toml:"db_max_open_conn_size"`
			DbMaxConnLifeTime int `toml:"db_max_conn_life_time"`
		} `toml:"sync"`
	} `toml:"database"`
	Redis struct {
		IPAddr   string `toml:"ip_addr"`
		Port     string `toml:"port"`
		Password string `toml:"password"`
		Sync     struct {
			MaxOpenConnSize int `toml:"max_open_conn_size"`
			MinOpenConnSize int `toml:"min_open_conn_size"`
			MaxConnLifeTime int `toml:"max_conn_life_time"`
			DialTimeout     int `toml:"dial_timeout"`
			PoolTimeout     int `toml:"pool_timeout"`
			IdleTimeout     int `toml:"idle_timeout"`
		} `toml:"sync"`
	} `toml:"redis"`
	Server struct {
		IPAddr   string `toml:"ip_addr"`
		Port     string `toml:"port"`
		Protocol string `toml:"protocol"`
	} `toml:"server"`
	Project struct {
		Mode               string   `toml:"mode"`
		UploadImageType    []string `toml:"upload_image_type"`
		UploadVideoType    []string `toml:"upload_video_type"`
		CosHeadPortraitDir string   `toml:"cos_head_portrait_dir"`
		CosImgDir          string   `toml:"cos_img_dir"`
		CosVideoDir        string   `toml:"cos_video_dir"`
		Auth               struct {
			TokenType        string `toml:"token_type"`
			TokenSignture    string `toml:"token_signture"`
			TokenTimeout     int    `toml:"token_timeout"`
			TokenSigntureKey string `toml:"token_signture_key"`
			WechatLogin      struct {
				AppID      string `toml:"app_id"`
				AppSercret string `toml:"app_sercret"`
			} `toml:"wechat_login"`
		} `toml:"auth"`
		Encrypt struct {
			RsaKeySize int `toml:"rsa_Key_size"`
		} `toml:"encrypt"`
		Cos struct {
			Appid              int    `toml:"appid"`
			SecretID           string `toml:"secret_id"`
			SecretKey          string `toml:"secret_key"`
			DurationSeconds    int    `toml:"duration_seconds"`
			PublicSourceBucket struct {
				Name   string `toml:"name"`
				Region string `toml:"region"`
			} `toml:"public_source_bucket"`
			PrivateSourceBucket struct {
				Name   string `toml:"name"`
				Region string `toml:"region"`
			} `toml:"private_source_bucket"`
		} `toml:"cos"`
	} `toml:"project"`
}

type BusinessConfig struct {
	CentRPCServer struct {
		IP   string `toml:"ip"`
		Port string `toml:"port"`
	} `toml:"cent_rpc_server"`
}

type IOWrite struct {
	Data []byte
}

func (i *IOWrite) Write(p []byte) (int, error) {
	i.Data = p
	return 0, nil
}

// 获得新的app.conf.toml
func GetNewAppConfig() (*AutoGenerated,error) {
	c,err := Config(&AutoGenerated{}).GetConfig()
	if err != nil {
		return nil, err
	}
	return c.(*AutoGenerated),nil
}

// 获得新的business.conf.toml
func GetNewBusConfig() (*BusinessConfig,error) {
	c,err := Config(&BusinessConfig{}).GetConfig()
	if err != nil {
		return nil, err
	}
	return c.(*BusinessConfig),nil
}


type Config interface {
	GetConfig() (interface{}, error)
	SetConfig(interface{}) error
	Unmarshal([]byte) (interface{}, error)
	Marshal() []byte
}

func (ag *AutoGenerated) GetConfig() (interface{}, error) {
	localPath, _ := os.Getwd()
	_, err := toml.DecodeFile(localPath+"/"+path.ConfFilePath+"/"+path.ConfRpcServerFileName, ag)
	if err != nil {
		return nil, errors.Wrap(err, ConfDecodeErr.Error())
	}
	return ag, err
}

func (ag *AutoGenerated) SetConfig(conf interface{}) error {
	localPath, _ := os.Getwd()
	result := conf.(*AutoGenerated)
	file, err := os.Create(localPath + "/" + path.ConfFilePath + "/" + path.ConfRpcServerFileName)
	if err != nil {
		return err
	}
	err = toml.NewEncoder(file).Encode(result)
	if err != nil {
		return errors.Wrap(err, ConfEncodeErr.Error())
	}
	return nil
}

// 返回github.com/pkg/errors包的自定义错误
func (ag *AutoGenerated) Unmarshal(bytes []byte) (interface{}, error) {
	err := toml.Unmarshal(bytes, ag)
	if err != nil {
		return nil, errors.Wrap(err, ConfUnmarshalErr.Error())
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

func (b *BusinessConfig) GetConfig() (interface{}, error) {
	localPath, _ := os.Getwd()
	_, err := toml.DecodeFile(localPath+"/"+path.ConfFilePath+"/"+path.ConfBusinessFileName, b)
	if err != nil {
		return nil, errors.Wrap(err, ConfDecodeErr.Error())
	}
	return b, err
}

func (b *BusinessConfig) SetConfig(conf interface{}) error {
	localPath, _ := os.Getwd()
	result := conf.(*BusinessConfig)
	file, err := os.Create(localPath + "/" + path.ConfFilePath + "/" + path.ConfBusinessFileName)
	if err != nil {
		return err
	}
	err = toml.NewEncoder(file).Encode(result)
	if err != nil {
		return errors.Wrap(err, ConfEncodeErr.Error())
	}
	return nil
}

func (b *BusinessConfig) Unmarshal(bytes []byte) (interface{}, error) {
	err := toml.Unmarshal(bytes, b)
	if err != nil {
		return nil, errors.Wrap(err, ConfUnmarshalErr.Error())
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

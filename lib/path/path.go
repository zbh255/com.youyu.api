package path

// 存储一些常用的项目path和默认设置
const (
	ConfDeVVersion        string = "dev"
	ConfProVersion        string = "pro"
	ConfFilePath          string = "/conf"
	ConfBusinessFileName  string = "business.conf.toml"
	ConfRpcServerFileName string = "app.conf.toml"

	// ConfBusinessRequestType 配置中心用来获取和写入配置的Config结构中的Type字段
	ConfBusinessRequestType string = "business_config"
	ConfRpcRequestType      string = "rpc_config"
	// LogGlobalPath 日志中心的日志文件路径
	LogGlobalPath string = "./log"
	// LogWebServerFileName gin web服务器日志文件
	LogWebServerFileName  string = "gin.log"
	LogDataRpcFileName    string = "data_rpc.log"
	LogConfigCentFileName string = "cent_rpc.log"
	LogBusinessFileName   string = "business.log"
	// InfoFileDefaultPath 信息文件的默认路径
	InfoFileDefaultPath string = "./dir"
	// ErrMsgJsonFileName 信息文件的默认名
	ErrMsgJsonFileName    string = "err_msg.json"
	RsaPublicKeyFileName  string = "pub_key.pem"
	RsaPrivateKeyFileName string = "pri_key.pem"
	// SigningKeyPrefix redis 用于区分存放签名密钥和token的key前缀
	SigningKeyPrefix string = "signing;key:"
	// TokenKeyPrefix redis 用于区分登录token的key前缀
	TokenKeyPrefix      string = "token;key:"
	TokenKeyWechatLogin string = "token;wechat_login;key:"
	TokenKeyWeiboLogin  string = "token;weibo_login;key:"
	// PubAndPriKeyPrefix redis 用于区分存储公私钥和token和签名的key
	PubAndPriKeyPrefix string = "pub.pri;x509;key:"
)

package path

// 存储一些常用的项目path和默认设置
const (
	ConfDeVVersion string = "dev"
	ConfProVersion string = "pro"
	ConfFilePath   string = "/conf"
	ConfBusinessFileName   string = "business.conf.toml"
	ConfRpcServerFileName string = "app.conf.toml"

	// 配置中心用来获取和写入配置的Config结构中的Type字段
	ConfBusinessRequestType string = "business_config"
	ConfRpcRequestType string = "rpc_config"
	// 日志中心的日志文件路径
	LogGlobalPath string = "./log"
	// 应用程序日志文件
	LogBusinessFileName string = "gin.log"
	LogDataRpcFileName string = "data_rpc.log"
	LogConfigCentFileName string = "cent_rpc.log"
)

package ecode

// 业务服务器错误号段[10000-19999]

var (
	// JsonParseError 程序解析参数失败
	JsonParseError = New(10001) // Json参数解析失败
	UrlParseError  = New(10002) // Url参数解析失败
	TomlParseError = New(10003) // Toml参数解析失败
	XmlParseError  = New(10004) // Xml参数解析失败
	ClientIdError  = New(10005) // 客户端uuid不合法
	EncodeError = New(10006) // 编码错误

	// SqlErr 外部程序出错或程序调用出错
	SqlErr         = New(10011) // sql不合法
	MySqlServerErr = New(10012) // Mysql服务器错误
	RedisServerErr = New(10013) // Redis服务器错误
	GrpcServerErr  = New(10014) // Grpc服务器错误

	// 文章模块
	AddArticleErr           = New(10201) // 添加文章失败
	DelArticleErr           = New(10202) // 删除文章失败
	UpdArticleErr           = New(10203) // 更新文章失败
	GetArticleErr           = New(10204) // 查询不到该文章
	AddArticleHotErr        = New(10205) // 添加文章热度失败
	AddArticleFabulousErr   = New(10206) // 添加文章点赞失败
	DelArticleFabulousErr   = New(10207) // 删除文章点赞失败
	GetArticleStatisticsErr = New(10208) // 获取文章文章的热度、点赞信息失败
	AddArticleCommentNumErr = New(10209) // 添加文章评论数失败

	// 广告模块
	GetAdvertisementErr = New(10301) // 获取广告失败
	AddAdvertisementErr = New(10302) // 添加广告失败
	UpdAdvertisementErr = New(10303) // 更新广告失败
	DelAdvertisementErr = New(10304) // 删除广告失败

	// 用户信息模块
	UserSignErr   = New(10401) //用户注册失败
	UserLoginErr   = New(10402) //用户登录失败
	UserDeleteErr = New(10403) // 用户删除失败

	// Tag 标签模块
	TagNameAlreadyExists = New(10501) // 标签名已经存在
	TagNameNotExists = New(10502) // 标签名不存在
	TagIdNotExists = New(10503) // 标签id不存在

	// 用户消息模块

	// 页面模块

	// 评论模块

	// 设置模块

	// 管理后台模块
)

package ecode

// All common ecode
var (
	OK = add(0) // 正确

	AppKeyInvalid           = add(-1)   // 应用程序不存在或已被封禁
	AccessKeyErr            = add(-2)   // Access Key错误
	SignCheckErr            = add(-3)   // API校验密匙错误
	MethodNoPermission      = add(-4)   // 调用方对该Method没有权限
	SecretKeyTimeout        = add(-5)   // 密钥已过期
	ParaMeterErr            = add(-6)   // 参数错误
	NoLogin                 = add(-101) // 账号未登录
	UserDisabled            = add(-102) // 账号被封停
	CaptchaErr              = add(-105) // 验证码错误
	UserInactive            = add(-106) // 账号未激活
	AppDenied               = add(-108) // 账号未激活
	MobileNoVerfiy          = add(-110) // 账号未激活
	CsrfNotMatchErr         = add(-111) // csrf 校验失败
	ServiceUpdate           = add(-112) // 系统升级中
	UserIDCheckInvalid      = add(-113) // 账号尚未实名认证
	UserIDCheckInvalidPhone = add(-114) // 请先绑定手机
	UserIDCheckInvalidCard  = add(-115) // 请先完成实名认证

	NotModified           = add(-304) // 木有改动
	TemporaryRedirect     = add(-307) // 撞车跳转
	RequestErr            = add(-400) // 请求错误
	Unauthorized          = add(-401) // 未认证
	AccessDenied          = add(-403) // 访问权限不足
	NothingFound          = add(-404) // 啥都木有
	ServerErr             = add(-500) // 服务器错误
	ServiceUnavailable    = add(-503) // 过载保护,服务暂不可用
	Deadline              = add(-504) // 服务调用超时
	LimitExceed           = add(-509) // 超出限制
	FileNotExists         = add(-616) // 上传文件不存在
	FileTooLarge          = add(-617) // 上传文件太大
	FileTypeErr           = add(-618) // 上传文件的类型错误
	FailedTooManyTimes    = add(-625) // 登录失败次数太多
	UserNotExist          = add(-626) // 用户不存在
	PasswordTooLeak       = add(-628) // 密码太弱
	UsernameOrPasswordErr = add(-629) // 用户名或密码错误
	UserLevelLow          = add(-650) // 用户等级太低
	UserDuplicate         = add(-652) // 重复的用户
	AccessTokenExpires    = add(-658) // Token 过期
	AccessTokenSignature  = add(-659) // 令牌签名验证失败
	AccessTokenErr        = add(-660) // 令牌不正确
	PasswordHashExpires   = add(-662) // 密码时间戳过期
	VcCodeTimeout         = add(-663) // 验证码已过期
	VcCodeError           = add(-664) // 验证码错误
)

package errors

import (
	"net/http"
)

/*
错误码设计
第一位表示错误级别, 1 为系统错误, 2 为普通错误
第二三四位表示服务模块代码
第五六位表示具体错误代码
*/

type Errno struct {
	Code     int    // 错误码
	Message  string // 展示给用户看的
	HttpCode int    // HTTP状态码
}

var (
	OK         = &Errno{Code: 0, Message: "OK", HttpCode: http.StatusOK}
	DelSignOK  = &Errno{Code: 1, Message: "DEL USER SIGN OK", HttpCode: http.StatusOK}
	UserSignOk = &Errno{Code: 2, Message: "USER SIGN OK", HttpCode: http.StatusOK}
	UploadOK   = &Errno{Code: 3, Message: "Upload OK", HttpCode: http.StatusOK}
	// 系统错误
	ErrUnKnown        = &Errno{Code: 100000, Message: "未知错误", HttpCode: http.StatusInternalServerError}
	ErrInternalServer = &Errno{Code: 100001, Message: "内部服务器错误", HttpCode: http.StatusInternalServerError}
	ErrParamConvert   = &Errno{Code: 100002, Message: "参数转换时发生错误", HttpCode: http.StatusInternalServerError}
	ErrDatabase       = &Errno{Code: 100003, Message: "数据库错误", HttpCode: http.StatusInternalServerError}
	ErrRedis          = &Errno{Code: 100004, Message: "Redis错误", HttpCode: http.StatusInternalServerError}

	// 模块通用错误
	ErrValidation      = &Errno{Code: 200001, Message: "参数校验失败", HttpCode: http.StatusForbidden}
	ErrBadRequest      = &Errno{Code: 200002, Message: "请求参数错误", HttpCode: http.StatusBadRequest}
	ErrGetTokenFail    = &Errno{Code: 200003, Message: "获取 token 失败", HttpCode: http.StatusForbidden}
	ErrTokenNotFound   = &Errno{Code: 200004, Message: "用户 token 不存在", HttpCode: http.StatusUnauthorized}
	ErrTokenExpire     = &Errno{Code: 200005, Message: "用户 token 过期", HttpCode: http.StatusForbidden}
	ErrTokenValidation = &Errno{Code: 200005, Message: "用户 token 无效", HttpCode: http.StatusForbidden}
	ErrTimeNoSwitch    = &Errno{Code: 200006, Message: "提交时间不正确", HttpCode: http.StatusForbidden}
	ErrJsonArgFailed   = &Errno{Code: 200007, Message: "提交的Json参数有误", HttpCode: http.StatusBadRequest}
	ErrRsaPublicKey    = &Errno{Code: 200008, Message: "用户的密钥不正确", HttpCode: http.StatusBadRequest}

	// User模块错误
	ErrUserNotFound            = &Errno{Code: 200104, Message: "用户不存在", HttpCode: http.StatusBadRequest}
	ErrPasswordIncorrect       = &Errno{Code: 200105, Message: "密码错误", HttpCode: http.StatusBadRequest}
	ErrUserRegisterAgain       = &Errno{Code: 200107, Message: "重复注册", HttpCode: http.StatusBadRequest}
	ErrUserNameExist           = &Errno{Code: 200108, Message: "用户名已存在", HttpCode: http.StatusBadRequest}
	ErrUsernameValidation      = &Errno{Code: 200109, Message: "用户名不合法", HttpCode: http.StatusBadRequest}
	ErrPasswordValidation      = &Errno{Code: 200110, Message: "密码不合法", HttpCode: http.StatusBadRequest}
	ErrUserSignNotFound        = &Errno{Code: 200111, Message: "该用户未登录", HttpCode: http.StatusBadRequest}
	ErrUserFriendRequest       = &Errno{Code: 200112, Message: "好友申请创建失败", HttpCode: http.StatusInternalServerError}
	ErrUserFriendRequestFailed = &Errno{Code: 200113, Message: "没有该用户的好友申请", HttpCode: http.StatusBadRequest}
	ErrUserFriendNotFound      = &Errno{Code: 200114, Message: "该好友不存在", HttpCode: http.StatusBadRequest}
	ErrUserPortrait            = &Errno{Code: 200115, Message: "该用户的头像路径为空", HttpCode: http.StatusBadRequest}
	// Group模块错误
	ErrGroupNotFound = &Errno{Code: 200201, Message: "分组不存在", HttpCode: http.StatusForbidden}
	ErrGroupExist    = &Errno{Code: 200202, Message: "分组已存在", HttpCode: http.StatusBadRequest}
	ErrGroupDefault  = &Errno{Code: 200203, Message: "默认分组不允许更改", HttpCode: http.StatusBadRequest}
	ErrGroupReName   = &Errno{Code: 200204, Message: "重命名的分组不能与存在的相同", HttpCode: http.StatusBadRequest}

	// 文件上传模块错误
	ErrUserUploadNotFound  = &Errno{Code: 200301, Message: "上传的文件为空", HttpCode: http.StatusBadRequest}
	ErrUserUploadExceedMax = &Errno{Code: 200302, Message: "上传的文件超过最大限制", HttpCode: http.StatusBadRequest}
	ErrUserUploadFormatNo  = &Errno{Code: 200303, Message: "上传的文件格式不正确", HttpCode: http.StatusBadRequest}
	ErrUserUploadUrlNot    = &Errno{Code: 200304, Message: "上传文件保存路径不存在", HttpCode: http.StatusBadRequest}
	ErrTencentCosLinkError = &Errno{Code: 200305, Message: "腾讯COS连接失败", HttpCode: http.StatusBadRequest}
	ErrTencentCosUploadNot = &Errno{Code: 200306, Message: "上传文件至COS失败", HttpCode: http.StatusBadRequest}
	ErrTencentCosNotFound  = &Errno{Code: 200307, Message: "COS中的文件不存在", HttpCode: http.StatusNotFound}

	// 游戏申请模块错误
	ErrUserInviteNotFound = &Errno{Code: 200401, Message: "该游戏申请不存在", HttpCode: http.StatusNotFound}
	ErrUserInviteNot      = &Errno{Code: 200402, Message: "没有查询到该用户的游戏申请", HttpCode: http.StatusNotFound}

	// 游戏模块错误
	ErrUserGameUnconfirmed = &Errno{Code: 200501, Message: "该用户未确认"}

	// 用户操作错误
	ErrDataBaseResultIsZero = &Errno{Code: 200601, Message: "未查询到该结果", HttpCode: http.StatusBadRequest}
)

func GetErrorsStruct(failed *Errno) *Errno {
	return failed
}

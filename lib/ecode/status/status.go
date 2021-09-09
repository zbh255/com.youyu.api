// Package status 从Grpc创建兼容业务代码的错误信息
package status

import (
	"com.youyu.api/lib/ecode"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
)

type Status struct {
	Code    int32
	Message string
}

func Error(code ecode.Code, message string) error {
	if code <= 10000 {
		c := math.Abs(float64(code))
		return status.Error(codes.Code(c), message)
	} else {
		return status.Error(codes.Code(code), message)
	}
}

// 为nil则填充正确信息
func FromError(err error) (*Status, bool) {
	if err == nil {
		return &Status{
			Code:    int32(ecode.OK),
			Message: ecode.OK.Message(),
		}, true
	}

	if st, bl := status.FromError(err); bl {
		if uint(st.Code()) > 10000 {
			return &Status{
				Code:    int32(st.Code()),
				Message: st.Message(),
			}, bl
		} else {
			return &Status{
				Code:    int32(st.Code() - st.Code()*2),
				Message: st.Message(),
			}, bl
		}
	} else {
		return nil, bl
	}
}

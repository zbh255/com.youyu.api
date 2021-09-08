// Package ecode 新的错误处理方式，比lib/errors更加完善
package ecode

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"sync/atomic"
)

// Codes 业务提示信息的接口
type Codes interface {
	// Error 有时会返回字符串形式的错误码
	// Error 报告错误
	Error() string
	// Code 返回包含的错误码
	Code() int
	// Message 返回错误包含的消息
	Message() string
}

type Code int

var (
	__message__ atomic.Value
	__codes__   = map[int]struct{}{}
)

// Register 提供给外部注册错误的函数
func Register(cam map[int]string) {
	__message__.Store(cam)
}

// New 业务提示信息的错误码
func New(e int) Code {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

// 添加基本错误信息码
func add(e int) Code {
	if _, ok := __codes__[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	__codes__[e] = struct{}{}
	return Code(e)
}

// GetRegisCodes 获得所有注册的错误码
func GetRegisCodes() []int {
	ints := make([]int, len(__codes__))
	ptr := 0
	for k, e := range __codes__ {
		if e == struct{}{} {
			ints[ptr] = k
			ptr++
		}
	}
	return ints
}

func (c Code) Error() string {
	return strconv.FormatInt(int64(c), 10)
}

func (c Code) Code() int {
	return int(c)
}

// Message 读取Register注册的错误代码对应的消息，消息不存在则返回空字符串
func (c Code) Message() string {
	if m, ok := __message__.Load().(map[int]string); ok {
		if msg, ok := m[c.Code()]; ok {
			return msg
		}
	}
	return c.Error()
}

// Equal 判断两个错误是否相同
func Equal(a, b Codes) bool {
	if a != nil && b != nil {
		return a.Code() == b.Code()
	} else {
		return false
	}
}

// Cause 将Err的根因错误转换为Codes接口,使其Erros方法可以获取错误信息
func Cause(e error) Codes {
	if e == nil {
		return OK
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		return ec
	}
	return String(e.Error())
}

func String(e string) Codes {
	if e == "" {
		return OK
	}
	// try error string
	i, err := strconv.Atoi(e)
	if err != nil {
		return ServerErr
	}
	return Code(i)
}

package log

import (
	"github.com/rs/zerolog"
	"io"
)

var Level = zerolog.ErrorLevel
var Logger zerolog.Logger

// 初始化日志的设置
func Init(w io.Writer)  {
	zerolog.SetGlobalLevel(Level)
	Logger = zerolog.New(w)
}



package log

import (
	"github.com/rs/zerolog"
)

//var Level = zerolog.ErrorLevel
//var Logger zerolog.Logger
//
//// Init old api
//// Init 初始化日志的设置
//func Init(w io.Writer) {
//	zerolog.SetGlobalLevel(Level)
//	Logger = zerolog.New(w)
//}

// Logger LoggerI new code
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(err error)
	Panic(err error)
}

type ZLogger struct {
	Level  zerolog.Level
	Logger zerolog.Logger
}

func (Z *ZLogger) Init() *ZLogger {
	zerolog.SetGlobalLevel(Z.Level)
	return Z
}

func (Z *ZLogger) Debug(msg string) {
	Z.Logger.Debug().Timestamp().Msg(msg)
}

func (Z *ZLogger) Info(msg string) {
	Z.Logger.Info().Timestamp().Msg(msg)
}

func (Z *ZLogger) Error(err error) {
	Z.Logger.Error().Timestamp().Msgf("%+v", err)
}

func (Z *ZLogger) Panic(err error) {
	Z.Logger.Panic().Timestamp().Err(err)
}

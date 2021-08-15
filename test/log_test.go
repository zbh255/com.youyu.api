package test

import (
	"com.youyu.api/app/rpc/client"
	"com.youyu.api/common/path"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	file, err := os.OpenFile("./log/gin.log", os.O_WRONLY|os.O_SYNC|os.O_APPEND, 0755)
	//file2,err := os.Open("./log/gin.log")
	if err != nil {
		t.Error(err)
	}
	var ioW io.Writer = os.Stdout
	ioW = file
	w := io.MultiWriter(ioW, os.Stdout)
	_, _ = w.Write([]byte("hello world"))
}

func TestZeroLog(t *testing.T) {
	file, _ := os.OpenFile(path.LogGlobalPath+"/"+path.LogConfigCentFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	logger := log.Output(io.MultiWriter(file, os.Stdout))
	logger.Error().Timestamp().Msg(fmt.Sprintf("%+v", errors.Wrap(errors.New("hello"), "world")))
	logger.Panic().Timestamp().Msg(fmt.Sprintf("%+v", errors.Wrap(errors.New("hello"), "world")))
}

func TestGinLog(t *testing.T) {
	eng := gin.New()
	gin.DefaultWriter = io.MultiWriter(&client.IOWTEST{})
	eng.Use(gin.Logger())
	_ = eng.Run()
}

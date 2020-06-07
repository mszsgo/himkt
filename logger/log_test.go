package logger

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

var log = Now("test")

func TestLog(t *testing.T) {
	log.WithField("method", "hm.test.query").Info("测试日志")
	// 不让程序瞬间退出
	ch := make(chan os.Signal)
	signal.Notify(ch)
	if _, ok := <-ch; ok {
		fmt.Println("Get the shutdown signal and start to shut down")
	}
}

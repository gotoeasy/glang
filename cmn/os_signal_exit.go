package cmn

import (
	"os"
	"os/signal"
	"syscall"
)

var exitFuncs []func()

func init() {
	go func() {
		osc := make(chan os.Signal, 1)
		signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
		<-osc
		Info("收到退出信号准备退出")
		for _, fnExit := range exitFuncs {
			fnExit()
		}
	}()
}

// 注册退出处理函数，在异步接收到SIGTERM或SIGINT信号时执行
func OnExit(fnExit func()) {
	exitFuncs = append(exitFuncs, fnExit)
}

// 同步等待到退出信号后退出
func WaitExit(fnExit func()) {
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	<-osc
	if fnExit != nil {
		fnExit()
	}
}

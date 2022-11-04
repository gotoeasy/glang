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

func OnExit(fnExit func()) {
	exitFuncs = append(exitFuncs, fnExit)
}

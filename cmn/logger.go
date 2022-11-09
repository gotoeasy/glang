package cmn

import (
	"fmt"
	"log"
)

var logLevel int
var glc *GLogCenterClient

func init() {
	glc = NewGLogCenterClient(&GlcOptions{
		Url:    GetEnvStr("GLC_API_URL", ""),
		System: GetEnvStr("GLC_SYSTEM", "default"),
		ApiKey: GetEnvStr("GLC_API_KEY", ""),
		Enable: GetEnvBool("GLC_ENABLE", false),
	})
}

// 设定日志级别（trace/debug/info/warn/error/fatal）
func SetLogLevel(level string) {
	if EqualsIngoreCase("TRACE", level) {
		logLevel = 0
	} else if EqualsIngoreCase("DEBUG", level) {
		logLevel = 1
	} else if EqualsIngoreCase("INFO", level) {
		logLevel = 2
	} else if EqualsIngoreCase("WARN", level) {
		logLevel = 3
	} else if EqualsIngoreCase("ERROR", level) {
		logLevel = 4
	} else if EqualsIngoreCase("FATAL", level) {
		logLevel = 5
	}
}

// 打印Trace级别日志
func Trace(v ...any) {
	if logLevel <= 0 {
		v = append([]any{"TRACE"}, v...)
		log.Println(v...)
		go glc.PostLog(fmt.Sprint(v...))
	}
}

// 打印Debug级别日志
func Debug(v ...any) {
	if logLevel <= 1 {
		v = append([]any{"DEBUG"}, v...)
		log.Println(v...)
		go glc.PostLog(fmt.Sprint(v...))
	}
}

// 打印Info级别日志
func Info(v ...any) {
	if logLevel <= 2 {
		v = append([]any{"INFO"}, v...)
		log.Println(v...)
		go glc.PostLog(fmt.Sprint(v...))
	}
}

// 打印Warn级别日志
func Warn(v ...any) {
	if logLevel <= 3 {
		v = append([]any{"WARN"}, v...)
		log.Println(v...)
		go glc.PostLog(fmt.Sprint(v...))
	}
}

// 打印Error级别日志
func Error(v ...any) {
	if logLevel <= 4 {
		v = append([]any{"ERROR"}, v...)
		log.Println(v...)
		go glc.PostLog(fmt.Sprint(v...))
	}
}

// 打印Fatal级别日志
func Fatal(v ...any) {
	if logLevel <= 5 {
		v = append([]any{"FATAL"}, v...)
		log.Println(v...)
		go glc.PostLog(fmt.Sprint(v...))
	}
}

// 打印Fatal级别日志，然后退出
func Fatalln(v ...any) {
	v = append([]any{"FATAL"}, v...)
	go glc.PostLog(fmt.Sprint(v...))
	log.Fatalln(v...)
}

// 打印日志
func Println(v ...any) {
	log.Println(v...)
	go glc.PostLog(fmt.Sprint(v...))
}

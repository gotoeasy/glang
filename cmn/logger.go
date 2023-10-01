package cmn

import (
	"fmt"
	"log"
)

var logLevel int

// 设定日志级别（trace/debug/info/warn/error/fatal）
func SetLogLevel(level string) {
	if EqualsIngoreCase("DEBUG", level) {
		logLevel = 1
	} else if EqualsIngoreCase("INFO", level) {
		logLevel = 2
	} else if EqualsIngoreCase("WARN", level) {
		logLevel = 3
	} else if EqualsIngoreCase("ERROR", level) {
		logLevel = 4
	}
}

// 打印Debug级别日志
func Debug(v ...any) {
	glc.Debug(v...)
	if logLevel <= 1 {
		log.Println(append([]any{"DEBUG"}, v...)...)
	}
}

// 打印Info级别日志
func Info(v ...any) {
	glc.Info(v...)
	if logLevel <= 2 {
		log.Println(append([]any{"INFO"}, v...)...)
	}
}

// 打印Warn级别日志
func Warn(v ...any) {
	glc.Warn(v...)
	if logLevel <= 3 {
		log.Println(append([]any{"WARN"}, v...)...)
	}
}

// 打印Error级别日志
func Error(v ...any) {
	glc.Error(v...)
	if logLevel <= 4 {
		log.Println(append([]any{"ERROR"}, v...)...)
	}
}

// 打印日志
func Println(v ...any) {
	glc.Println(fmt.Sprint(v...))
	log.Println(v...)
}

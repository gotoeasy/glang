package cmn

import (
	"fmt"
	"log"
)

var logLevel int
var glcLogLevel int
var glc *GLogCenterClient

func init() {
	glc = NewGLogCenterClient(&GlcOptions{
		Url:    GetEnvStr("GLC_API_URL", "http://glc.nnn.com/glc/v1/log/add"),
		System: GetEnvStr("GLC_SYSTEM", "glang/cmn"),
		ApiKey: GetEnvStr("GLC_API_KEY", ""),
		Enable: GetEnvBool("GLC_ENABLE", false), // 默认关闭GLC
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

	glcLevel := GetEnvStr("GLC_LOG_LEVEL", "trace")
	if EqualsIngoreCase("TRACE", glcLevel) {
		glcLogLevel = 0
	} else if EqualsIngoreCase("DEBUG", glcLevel) {
		glcLogLevel = 1
	} else if EqualsIngoreCase("INFO", glcLevel) {
		glcLogLevel = 2
	} else if EqualsIngoreCase("WARN", glcLevel) {
		glcLogLevel = 3
	} else if EqualsIngoreCase("ERROR", glcLevel) {
		glcLogLevel = 4
	} else if EqualsIngoreCase("FATAL", glcLevel) {
		glcLogLevel = 5
	}

}

// 打印Trace级别日志
func Trace(v ...any) {
	if glc.enable && glcLogLevel <= 0 {
		go glc.PostLog("TRACE " + fmt.Sprint(v...))
	}
	if logLevel <= 0 {
		log.Println(append([]any{"TRACE"}, v...)...)
	}
}

// 打印Debug级别日志
func Debug(v ...any) {
	if glc.enable && glcLogLevel <= 1 {
		go glc.PostLog("DEBUG " + fmt.Sprint(v...))
	}
	if logLevel <= 1 {
		log.Println(append([]any{"DEBUG"}, v...)...)
	}
}

// 打印Info级别日志
func Info(v ...any) {
	if glc.enable && glcLogLevel <= 2 {
		go glc.PostLog("INFO " + fmt.Sprint(v...))
	}
	if logLevel <= 2 {
		log.Println(append([]any{"INFO"}, v...)...)
	}
}

// 打印Warn级别日志
func Warn(v ...any) {
	if glc.enable && glcLogLevel <= 3 {
		go glc.PostLog("WARN " + fmt.Sprint(v...))
	}
	if logLevel <= 3 {
		log.Println(append([]any{"WARN"}, v...)...)
	}
}

// 打印Error级别日志
func Error(v ...any) {
	if glc.enable && glcLogLevel <= 4 {
		go glc.PostLog("ERROR " + fmt.Sprint(v...))
	}
	if logLevel <= 4 {
		log.Println(append([]any{"ERROR"}, v...)...)
	}
}

// 打印Fatal级别日志
func Fatal(v ...any) {
	if glc.enable && glcLogLevel <= 5 {
		go glc.PostLog("FATAL " + fmt.Sprint(v...))
	}
	if logLevel <= 5 {
		log.Println(append([]any{"FATAL"}, v...)...)
	}
}

// 打印Fatal级别日志，然后退出
func Fatalln(v ...any) {
	if glc.enable {
		go glc.PostLog("FATAL " + fmt.Sprint(v...))
	}
	log.Fatalln(append([]any{"FATAL"}, v...)...)

}

// 打印日志
func Println(v ...any) {
	if glc.enable {
		go glc.PostLog(fmt.Sprint(v...))
	}
	log.Println(v...)
}

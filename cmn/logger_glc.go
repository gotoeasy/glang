package cmn

import (
	"fmt"
	"strings"
)

// 日志中心客户端结构体
type GLogCenterClient struct {
	url      string
	system   string
	apiKey   string
	enable   bool
	logLevel int
	logChan  chan string // 用chan控制日志发送顺序
}

// 日志中心选项
type GlcOptions struct {
	Url      string
	System   string
	ApiKey   string
	Enable   bool
	LogLevel string
}

// 创建日志中心客户端对象
func NewGLogCenterClient(o *GlcOptions) *GLogCenterClient {
	if o == nil {
		o = &GlcOptions{}
	}

	glc := &GLogCenterClient{
		url:     o.Url,
		system:  o.System,
		apiKey:  o.ApiKey,
		enable:  o.Enable,
		logChan: make(chan string, 2048),
	}

	if EqualsIngoreCase("DEBUG", o.LogLevel) {
		glc.logLevel = 1
	} else if EqualsIngoreCase("INFO", o.LogLevel) {
		glc.logLevel = 2
	} else if EqualsIngoreCase("WARN", o.LogLevel) {
		glc.logLevel = 3
	} else if EqualsIngoreCase("ERROR", o.LogLevel) {
		glc.logLevel = 4
	} else if EqualsIngoreCase("FATAL", o.LogLevel) {
		logLevel = 5
	}

	go func() {
		for {
			text := <-glc.logChan
			FasthttpPostJson(glc.url, text, glc.apiKey)
		}
	}()

	return glc
}

// 发送Trace级别日志到日志中心
func (g *GLogCenterClient) Trace(v ...any) {
	if glc.enable && glc.logLevel <= 0 {
		g.SentLog("TRACE " + fmt.Sprint(v...))
	}
}

// 发送Debug级别日志到日志中心
func (g *GLogCenterClient) Debug(v ...any) {
	if glc.enable && glc.logLevel <= 1 {
		g.SentLog("DEBUG " + fmt.Sprint(v...))
	}
}

// 发送Info级别日志到日志中心
func (g *GLogCenterClient) Info(v ...any) {
	if glc.enable && glc.logLevel <= 2 {
		g.SentLog("INFO " + fmt.Sprint(v...))
	}
}

// 发送Warn级别日志到日志中心
func (g *GLogCenterClient) Warn(v ...any) {
	if glc.enable && glc.logLevel <= 3 {
		g.SentLog("WARN " + fmt.Sprint(v...))
	}
}

// 发送Error级别日志到日志中心
func (g *GLogCenterClient) Error(v ...any) {
	if glc.enable && glc.logLevel <= 4 {
		g.SentLog("ERROR " + fmt.Sprint(v...))
	}
}

// 发送Fatal级别日志到日志中心
func (g *GLogCenterClient) Fatal(v ...any) {
	if glc.enable && glc.logLevel <= 5 {
		g.SentLog("FATAL " + fmt.Sprint(v...))
	}
}

// 发送指定级别日志到日志中心
func (g *GLogCenterClient) SentLog(text string) {
	if IsBlank(text) {
		return
	}
	var data strings.Builder
	data.WriteString("{")
	data.WriteString(`"system":"` + g.encodeGlcJsonValue(g.system) + `"`)
	data.WriteString(`,"date":"` + FormatSystemDate(FMT_YYYY_MM_DD_HH_MM_SS_SSS) + `"`)
	data.WriteString(`,"text":"` + g.encodeGlcJsonValue(text) + `"`)
	data.WriteString("}")

	g.logChan <- data.String()
}

func (g *GLogCenterClient) encodeGlcJsonValue(v string) string {
	v = ReplaceAll(v, `"`, `\"`)
	v = ReplaceAll(v, "\t", "\\\\t")
	v = ReplaceAll(v, "\r", "\\\\r")
	v = ReplaceAll(v, "\n", "\\\\n")
	return v
}

package cmn

import (
	"encoding/json"
	"fmt"
)

// 日志接口数据结构体
type LogDataModel struct {
	Text       string `json:"text,omitempty"`       // 【必须】日志内容，多行时仅为首行，直接显示用，是全文检索对象
	Date       string `json:"date,omitempty"`       // 日期（格式YYYY-MM-DD HH:MM:SS.SSS）
	System     string `json:"system,omitempty"`     // 系统名
	ServerName string `json:"servername,omitempty"` // 服务器名
	ServerIp   string `json:"serverip,omitempty"`   // 服务器IP
	ClientIp   string `json:"clientip,omitempty"`   // 客户端IP
	TraceId    string `json:"traceid,omitempty"`    // 跟踪ID
	LogType    string `json:"logtype,omitempty"`    // 日志类型（1:登录日志、2:操作日志）
	LogLevel   string `json:"loglevel,omitempty"`   // 日志级别
	User       string `json:"user,omitempty"`       // 用户
	Module     string `json:"module,omitempty"`     // 模块
	Operation  string `json:"action,omitempty"`     // 操作
}

func (d *LogDataModel) ToJson() string {
	bt, _ := json.Marshal(d)
	return BytesToString(bt)
}

// 日志中心客户端结构体
//
// 日志中心见 https://github.com/gotoeasy/glogcenter
type GLogCenterClient struct {
	apiUrl   string
	system   string
	apiKey   string
	enable   bool
	logLevel int
	logChan  chan string // 用chan控制日志发送顺序
}

// 日志中心选项
type GlcOptions struct {
	ApiUrl   string // 日志中心的添加日志接口地址
	System   string // 系统名（对应日志中心检索页面的分类栏）
	ApiKey   string // 日志中心的ApiKey
	Enable   bool   // 是否开启发送到日志中心
	LogLevel string // 日志级别（trace/debug/info/warn/error/fatal）
}

var glc *GLogCenterClient

// 按环境编配配置初始化glc对象，方便开箱即用，外部使用时可通过SetLogCenterClient重新设定
func init() {
	SetLogCenterClient(NewGLogCenterClient(&GlcOptions{
		ApiUrl:   GetEnvStr("GLC_API_URL", ""),
		System:   GetEnvStr("GLC_SYSTEM", "glang"),
		ApiKey:   GetEnvStr("GLC_API_KEY", ""),
		Enable:   GetEnvBool("GLC_ENABLE", false),
		LogLevel: GetEnvStr("GLC_LOG_LEVEL", "debug"),
	}))
}

// 创建日志中心客户端对象
func NewGLogCenterClient(o *GlcOptions) *GLogCenterClient {
	if o == nil {
		o = &GlcOptions{}
	}

	glc := &GLogCenterClient{
		apiUrl:  o.ApiUrl,
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
			FasthttpPostJson(glc.apiUrl, text, glc.apiKey)
		}
	}()

	return glc
}

// 设定GLC日志中心客户端
func SetLogCenterClient(glcClient *GLogCenterClient) {
	glc = glcClient
}

// 发送Trace级别日志到日志中心
func (g *GLogCenterClient) Trace(v ...any) {
	if glc.enable && glc.logLevel <= 0 {
		g.print(g.system, "TRACE", "TRACE "+fmt.Sprint(v...))
	}
}

// 发送指定系统名的Trace级别日志到日志中心
func (g *GLogCenterClient) TraceSys(system string, v ...any) {
	if glc.enable && glc.logLevel <= 0 {
		g.print(g.system, "TRACE", "TRACE "+fmt.Sprint(v...))
	}
}

// 发送Debug级别日志到日志中心
func (g *GLogCenterClient) Debug(v ...any) {
	if glc.enable && glc.logLevel <= 1 {
		g.print(g.system, "DEBUG", "DEBUG "+fmt.Sprint(v...))
	}
}

// 发送Debug级别日志到日志中心
func (g *GLogCenterClient) DebugSys(system string, v ...any) {
	if glc.enable && glc.logLevel <= 1 {
		g.print(g.system, "DEBUG", "DEBUG "+fmt.Sprint(v...))
	}
}

// 发送Info级别日志到日志中心
func (g *GLogCenterClient) Info(v ...any) {
	if glc.enable && glc.logLevel <= 2 {
		g.print(g.system, "INFO", "INFO "+fmt.Sprint(v...))
	}
}

// 发送指定系统名的Info级别日志到日志中心
func (g *GLogCenterClient) InfoSys(system string, v ...any) {
	if glc.enable && glc.logLevel <= 2 {
		g.print(system, "INFO", "INFO "+fmt.Sprint(v...))
	}
}

// 发送Warn级别日志到日志中心
func (g *GLogCenterClient) Warn(v ...any) {
	if glc.enable && glc.logLevel <= 3 {
		g.print(g.system, "WARN", "WARN "+fmt.Sprint(v...))
	}
}

// 发送指定系统名的Warn级别日志到日志中心
func (g *GLogCenterClient) WarnSys(system string, v ...any) {
	if glc.enable && glc.logLevel <= 3 {
		g.print(system, "WARN", "WARN "+fmt.Sprint(v...))
	}
}

// 发送Error级别日志到日志中心
func (g *GLogCenterClient) Error(v ...any) {
	if glc.enable && glc.logLevel <= 4 {
		g.print(g.system, "ERROR", "ERROR "+fmt.Sprint(v...))
	}
}

// 发送指定系统名的Error级别日志到日志中心
func (g *GLogCenterClient) ErrorSys(system string, v ...any) {
	if glc.enable && glc.logLevel <= 4 {
		g.print(system, "ERROR", "ERROR "+fmt.Sprint(v...))
	}
}

// 发送Fatal级别日志到日志中心
func (g *GLogCenterClient) Fatal(v ...any) {
	if glc.enable && glc.logLevel <= 5 {
		g.print(g.system, "FATAL", "FATAL "+fmt.Sprint(v...))
	}
}

// 发送指定系统名的Fatal级别日志到日志中心
func (g *GLogCenterClient) FatalSys(system string, v ...any) {
	if glc.enable && glc.logLevel <= 5 {
		g.print(system, "FATAL", "FATAL "+fmt.Sprint(v...))
	}
}

// 发送日志到日志中心
func (g *GLogCenterClient) Println(text string) {
	g.print(g.system, "INFO", text)
}

func (g *GLogCenterClient) print(system string, logLevel string, text string) {
	if IsBlank(text) {
		return
	}

	data := new(LogDataModel)
	data.System = system
	data.Date = FormatSystemDate(FMT_YYYY_MM_DD_HH_MM_SS_SSS)
	data.Text = text
	data.ServerIp = GetLocalIp()
	data.ServerName = GetLocalHostName()
	data.LogLevel = logLevel

	g.logChan <- data.ToJson()
}

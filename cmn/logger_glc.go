package cmn

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

// 日志接口数据结构体
type GlcData struct {
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

var ldmType reflect.Type = reflect.TypeOf(&GlcData{})

func (d *GlcData) ToJson() string {
	bt, _ := json.Marshal(d)
	return BytesToString(bt)
}

// 日志中心客户端结构体
//
// 日志中心见 https://github.com/gotoeasy/glogcenter
type GlcClient struct {
	apiUrl   string
	system   string
	apiKey   string
	enable   bool
	stop     bool
	logLevel int
	logChan  chan *GlcData // 用chan控制日志发送顺序
}

// 日志中心选项
type GlcOptions struct {
	ApiUrl   string // 日志中心的添加日志接口地址，默认取环境变量GLC_API_URL
	System   string // 系统名（对应日志中心检索页面的分类栏），默认取环境变量GLC_API_URL，未设定时default
	ApiKey   string // 日志中心的ApiKey，默认取环境变量GLC_API_URL
	Enable   bool   // 是否开启发送到日志中心，默认取环境变量GLC_API_URL，未设定时false
	LogLevel string // 能输出的日志级别（DEBUG/INFO/WARN/ERROR），默认取环境变量GLC_API_URL，未设定时DEBUG
}

var glc *GlcClient

// 创建日志中心客户端对象
func NewGlcClient(o *GlcOptions) *GlcClient {
	if o == nil {
		// 按环境编配配置初始化glc对象
		o = &GlcOptions{
			ApiUrl:   GetEnvStr("GLC_API_URL", ""),
			System:   GetEnvStr("GLC_SYSTEM", "default"),
			ApiKey:   GetEnvStr("GLC_API_KEY", ""),
			Enable:   GetEnvBool("GLC_ENABLE", false),
			LogLevel: GetEnvStr("GLC_LOG_LEVEL", "DEBUG"),
		}
	} else {
		if o.ApiUrl == "" {
			o.ApiUrl = GetEnvStr("GLC_API_URL", "")
		}
		if o.System == "" {
			o.System = GetEnvStr("GLC_SYSTEM", "default")
		}
		if o.ApiKey == "" {
			o.ApiKey = GetEnvStr("GLC_API_KEY", "")
		}
	}

	glc := &GlcClient{
		apiUrl:  o.ApiUrl,
		system:  o.System,
		apiKey:  o.ApiKey,
		enable:  o.Enable,
		logChan: make(chan *GlcData, 1024),
	}

	if EqualsIngoreCase("DEBUG", o.LogLevel) {
		glc.logLevel = 1
	} else if EqualsIngoreCase("INFO", o.LogLevel) {
		glc.logLevel = 2
	} else if EqualsIngoreCase("WARN", o.LogLevel) {
		glc.logLevel = 3
	} else if EqualsIngoreCase("ERROR", o.LogLevel) {
		glc.logLevel = 4
	}

	go func() {
		for {
			ldm := <-glc.logChan
			FasthttpPostJson(glc.apiUrl, ldm.ToJson(), glc.apiKey)
		}
	}()

	return glc
}

// 设定GLC日志中心客户端
func SetGlcClient(glcClient *GlcClient) {
	glc = glcClient
}

// 发送Debug级别日志到日志中心
func (g *GlcClient) Debug(v ...any) {
	params, ldm := logParams(v...)
	if g == nil {
		log.Println(append([]any{"DEBUG"}, params...)...)
	} else if g.logLevel <= 1 {
		// 控制台日志
		log.Println(append([]any{"DEBUG"}, params...)...)
		// GLC日志
		if !g.stop && g.enable {
			if ldm == nil {
				ldm = &GlcData{}
			}
			ldm.LogLevel = "DEBUG"
			g.print(params, ldm)
		}
	}
}

// 发送Info级别日志到日志中心
func (g *GlcClient) Info(v ...any) {
	params, ldm := logParams(v...)
	if glc == nil {
		log.Println(append([]any{"INFO"}, params...)...)
	} else if glc.logLevel <= 2 {
		// 控制台日志
		log.Println(append([]any{"INFO"}, params...)...)
		// GLC日志
		if !g.stop && g.enable {
			if ldm == nil {
				ldm = &GlcData{}
			}
			ldm.LogLevel = "INFO"
			g.print(params, ldm)
		}
	}
}

// 发送Warn级别日志到日志中心
func (g *GlcClient) Warn(v ...any) {
	params, ldm := logParams(v...)
	if glc == nil {
		log.Println(append([]any{"WARN"}, params...)...)
	} else if glc.logLevel <= 3 {
		// 控制台日志
		log.Println(append([]any{"WARN"}, params...)...)
		// GLC日志
		if !g.stop && g.enable {
			if ldm == nil {
				ldm = &GlcData{}
			}
			ldm.LogLevel = "WARN"
			g.print(params, ldm)
		}
	}
}

// 发送Error级别日志到日志中心
func (g *GlcClient) Error(v ...any) {
	params, ldm := logParams(v...)
	if glc == nil {
		log.Println(append([]any{"ERROR"}, params...)...)
	} else if glc.logLevel <= 4 {
		// 控制台日志
		log.Println(append([]any{"ERROR"}, params...)...)
		// GLC日志
		if !g.stop && g.enable {
			if ldm == nil {
				ldm = &GlcData{}
			}
			ldm.LogLevel = "ERROR"
			g.print(params, ldm)
		}
	}
}

func logParams(v ...any) ([]any, *GlcData) {
	var ary []any
	var ldm *GlcData
	for i := 0; i < len(v); i++ {
		if v[i] != nil {
			if reflect.TypeOf(v[i]) == ldmType {
				ldm = v[i].(*GlcData)
			} else {
				ary = append(ary, v[i])
			}
		}
	}
	return ary, ldm
}

func (g *GlcClient) print(params []any, ldm *GlcData) {
	// 日志参数优先
	if len(params) > 0 {
		ldm.Text = fmt.Sprint(params...)
	}

	// 其他字段检查补填
	if ldm.System == "" {
		ldm.System = g.system
	}
	ldm.Date = Now()
	if ldm.ServerIp == "" {
		ldm.ServerIp = GetPreferredLocalIPv4()
	}
	if ldm.ServerName == "" {
		ldm.ServerName = GetLocalHostName()
	}
	if ldm.TraceId == "" {
		ldm.TraceId = HashString(ULID())
	}

	g.logChan <- ldm
}

// 停止接收新的日志并等待日志全部输出完成
func (g *GlcClient) WaitFinish() {
	g.stop = true
	for {
		if len(g.logChan) <= 0 {
			break
		}
		time.Sleep(time.Millisecond * 10)
	}
}

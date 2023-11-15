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
	TraceId    string `json:"traceid,omitempty"`    // 跟踪码
	LogLevel   string `json:"loglevel,omitempty"`   // 日志级别
	User       string `json:"user,omitempty"`       // 用户
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
	opt      *GlcOptions
	stop     bool
	busy     bool
	logLevel int
	logChan  chan *GlcData // 用chan控制日志发送顺序
}

// 日志中心选项
type GlcOptions struct {
	ApiUrl            string // 日志中心的添加日志接口地址，默认取环境变量GLC_API_URL
	System            string // 系统名（对应日志中心检索页面的分类栏），默认取环境变量GLC_SYSTEM，未设定时default
	ApiKey            string // 日志中心的ApiKey，默认取环境变量GLC_API_KEY，未设定时X-GLC-AUTH:glogcenter
	Enable            bool   // 是否开启发送到日志中心，默认取环境变量GLC_ENABLE，未设定时false
	DisableConsoleLog bool   // 是否禁止打印控制台日志，默认取环境变量GLC_DISABLE_CONSOLE_LOG，默认false
	LogLevel          string // 能输出的日志级别（DEBUG/INFO/WARN/ERROR），默认取环境变量GLC_LOG_LEVEL，未设定时DEBUG
	ServerName        string // 服务器名
	ServerIp          string // 服务器IP
	ClientIp          string // 客户端IP
	AddCityToIp       bool   // 是否添加城市信息到IP前面，默认取环境变量GLC_ADD_CITY_TO_IP，默认false
	PrintSrcLine      bool   // 是否添加打印调用的文件行号
}

var _glc *GlcClient

func init() {
	if GetEnvBool("GLC_ENABLE", false) {
		_glc = NewGlcClient(nil) // 使用环境变量配置初始化
	}
}

// 创建日志中心客户端对象
func NewGlcClient(o *GlcOptions) *GlcClient {
	if o == nil {
		// 按环境编配配置初始化glc对象
		o = &GlcOptions{
			ApiUrl:            GetEnvStr("GLC_API_URL", ""),
			System:            GetEnvStr("GLC_SYSTEM", "default"),
			ApiKey:            GetEnvStr("GLC_API_KEY", "X-GLC-AUTH:glogcenter"),
			Enable:            GetEnvBool("GLC_ENABLE", false),
			DisableConsoleLog: GetEnvBool("GLC_DISABLE_CONSOLE_LOG", false),
			LogLevel:          GetEnvStr("GLC_LOG_LEVEL", "DEBUG"),
			AddCityToIp:       GetEnvBool("GLC_ADD_CITY_TO_IP", false),
			PrintSrcLine:      GetEnvBool("GLC_PRINT_SRC_LINE", false),
		}
	} else {
		if o.ApiUrl == "" {
			o.ApiUrl = GetEnvStr("GLC_API_URL", "")
		}
		if o.System == "" {
			o.System = GetEnvStr("GLC_SYSTEM", "default")
		}
		if o.ApiKey == "" {
			o.ApiKey = GetEnvStr("GLC_API_KEY", "X-GLC-AUTH:glogcenter")
		}
	}

	glc := &GlcClient{
		opt:     o,
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
			gd := <-glc.logChan
			FasthttpPostJson(glc.opt.ApiUrl, gd.ToJson(), glc.opt.ApiKey)
			if len(glc.logChan) <= 0 {
				glc.busy = false
			}
		}
	}()

	return glc
}

// 设定GLC日志中心客户端
func SetGlcClient(glcClient *GlcClient) {
	_glc = glcClient
}

// 发送Debug级别日志到日志中心
func (g *GlcClient) Debug(v ...any) {
	params, ldm := logParams(v...)
	glcPrint(g, "DEBUG", params, ldm)
}

// 发送Info级别日志到日志中心
func (g *GlcClient) Info(v ...any) {
	params, ldm := logParams(v...)
	glcPrint(g, "INFO", params, ldm)
}

// 发送Warn级别日志到日志中心
func (g *GlcClient) Warn(v ...any) {
	params, ldm := logParams(v...)
	glcPrint(g, "WARN", params, ldm)
}

// 发送Error级别日志到日志中心
func (g *GlcClient) Error(v ...any) {
	params, ldm := logParams(v...)
	glcPrint(g, "ERROR", params, ldm)
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

func glcPrint(g *GlcClient, level string, params []any, ldm *GlcData) {

	if g == nil || !g.opt.DisableConsoleLog {
		log.Println(append([]any{level}, params...)...) // 控制台日志
	}
	if g == nil || g.stop || !g.opt.Enable || g.opt.ApiUrl == "" {
		return
	}

	glcData := &GlcData{
		Date:       Now(),
		System:     g.opt.System,
		ServerName: g.opt.ServerName,
		ServerIp:   g.opt.ServerIp,
		ClientIp:   g.opt.ClientIp,
		LogLevel:   level,
	}

	if len(params) <= 0 {
		if ldm != nil {
			glcData.Text = ldm.Text
		}
	} else {
		glcData.Text = fmt.Sprint(params...) // 日志参数优先
	}

	// 其他字段检查补填
	if ldm != nil {
		if ldm.System != "" {
			glcData.System = ldm.System
		}
		if ldm.ServerName != "" {
			glcData.ServerName = ldm.ServerName
		}
		if ldm.ServerIp != "" {
			glcData.ServerIp = ldm.ServerIp
		}
		if ldm.ClientIp != "" {
			glcData.ClientIp = ldm.ClientIp
		}
		if ldm.TraceId != "" {
			glcData.TraceId = ldm.TraceId
		}
	}
	if glcData.TraceId == "" {
		glcData.TraceId = HashString(ULID())
	}
	if g.opt.AddCityToIp {
		glcData.ServerIp = GetCityIp(glcData.ServerIp)
		glcData.ClientIp = GetCityIp(glcData.ClientIp)
	}

	g.busy = true
	g.logChan <- glcData
}

// 停止接收新的日志并等待日志全部输出完成
func (g *GlcClient) WaitFinish() {
	if g != nil {
		g.stop = true
		for {
			if !g.busy {
				break
			}
			time.Sleep(time.Millisecond * 10)
		}
	}
}

func GetGlcLatestVersion(url string) string {
	// 返回样例 {"version":"v0.12.0"}
	bts, err := HttpGetJson(url, "Glc:"+HashString(GetLocalHostName()+"一个端点有区别的固定值"))
	if err == nil {
		var data struct {
			Version string `json:"version,omitempty"`
		}
		if err := json.Unmarshal(bts, &data); err == nil {
			return data.Version
		}
	}
	return ""
}

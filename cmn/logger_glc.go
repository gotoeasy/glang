package cmn

import (
	"strings"
)

// 日志中心客户端结构体
type GLogCenterClient struct {
	url    string
	system string
	apiKey string
	enable bool
}

// 日志中心选项
type GlcOptions struct {
	Url    string
	System string
	ApiKey string
	Enable bool
}

// 创建日志中心客户端对象
func NewGLogCenterClient(o *GlcOptions) *GLogCenterClient {
	if o == nil {
		o = &GlcOptions{}
	}
	return &GLogCenterClient{
		url:    o.Url,
		system: o.System,
		apiKey: o.ApiKey,
		enable: o.Enable,
	}
}

// 发送日志到日志中心
func (g *GLogCenterClient) PostLog(text string) {
	text = Trim(text)
	if text == "" {
		return
	}

	var data strings.Builder
	data.WriteString("{")
	data.WriteString(`"system":"` + g.encodeGlcJsonValue(g.system) + `"`)
	data.WriteString(`,"date":"` + GetYyyymmddHHMMSS() + `"`)
	data.WriteString(`,"text":"` + g.encodeGlcJsonValue(text) + `"`)
	data.WriteString("}")

	FasthttpPostJson(g.url, data.String(), g.apiKey)
}

func (g *GLogCenterClient) encodeGlcJsonValue(v string) string {
	v = ReplaceAll(v, `"`, `\"`)
	v = ReplaceAll(v, "\t", "\\\\t")
	v = ReplaceAll(v, "\r", "\\\\r")
	v = ReplaceAll(v, "\n", "\\\\n")
	return v
}

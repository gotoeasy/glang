package cmn

import "strings"

// 日志中心客户端结构体
type GLogCenterClient struct {
	opt *GlcOptions
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
	return &GLogCenterClient{opt: o}
}

// 发送日志到日志中心
func (g *GLogCenterClient) PostLog(text string) {
	if !g.opt.Enable {
		return
	}

	text = Trim(text)
	if text == "" {
		return
	}

	var data strings.Builder
	data.WriteString("{")
	data.WriteString(`"system":"` + g.encodeGlcJsonValue(g.opt.System) + `"`)
	data.WriteString(`,"date":"` + GetYyyymmddHHMMSS() + `"`)
	data.WriteString(`,"text":"` + g.encodeGlcJsonValue(text) + `"`)
	data.WriteString("}")

	FasthttpPostJson(g.opt.Url, data.String(), g.opt.ApiKey)
}

func (g *GLogCenterClient) encodeGlcJsonValue(v string) string {
	v = ReplaceAll(v, `"`, `\"`)
	v = ReplaceAll(v, "\t", "\\\\t")
	v = ReplaceAll(v, "\r", "\\\\r")
	v = ReplaceAll(v, "\n", "\\\\n")
	return v
}

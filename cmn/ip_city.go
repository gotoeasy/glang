package cmn

import (
	"encoding/json"
	"time"
)

type ipCityResult struct {
	IP          string `json:"ip,omitempty"`
	Pro         string `json:"pro,omitempty"`  // 省
	City        string `json:"city,omitempty"` // 市
	ProCode     string `json:"proCode,omitempty"`
	CityCode    string `json:"cityCode,omitempty"`
	Region      string `json:"region,omitempty"`
	RegionCode  string `json:"regionCode,omitempty"`
	Addr        string `json:"addr,omitempty"`
	RegionNames string `json:"regionNames,omitempty"`
	Err         string `json:"err,omitempty"`
}

// IPInfoResponse 接口完整响应体
type ipInfoResponse struct {
	Ret  int      `json:"ret,omitempty"`
	Data ipDetail `json:"data,omitempty"`
	Qt   int      `json:"qt,omitempty"`
}

// IPDetail IP 详细信息（英文命名）
type ipDetail struct {
	IP            string `json:"ip,omitempty"`              // IP地址
	Country       string `json:"country,omitempty"`         // 国家
	CountryCode   string `json:"country_code,omitempty"`    // 国家代码
	Province      string `json:"prov,omitempty"`            // 省份
	City          string `json:"city,omitempty"`            // 城市
	CityCode      string `json:"city_code,omitempty"`       // 城市代码
	CityShortCode string `json:"city_short_code,omitempty"` // 城市短代码
	Area          string `json:"area,omitempty"`            // 区域/区县
	PostCode      string `json:"post_code,omitempty"`       // 邮编
	AreaCode      string `json:"area_code,omitempty"`       // 区号
	ISP           string `json:"isp,omitempty"`             // 运营商
	Longitude     string `json:"lng,omitempty"`             // 经度
	Latitude      string `json:"lat,omitempty"`             // 纬度
	LongIP        uint32 `json:"long_ip,omitempty"`         // 整型IP
	BigArea       string `json:"big_area,omitempty"`        // 大区
}

var _ipCache *LruCache

// 获取ip地址信息含ip
func GetCityIp(ip string) string {
	if ip == "" {
		return ""
	}
	ct := GetCityByIp(ip)
	if ct != "" {
		return ct + " " + ip
	}
	return ip
}

// 获取ip地址信息不含ip
func GetCityByIp(ip string) string {
	rs := GetCityByIp_ip9(ip)
	if rs == "" {
		rs = GetCityByIp_pconline(ip)
	}
	return rs
}

// 获取ip地址信息不含ip
func GetCityByIp_pconline(ip string) string {
	if _ipCache == nil {
		_ipCache = NewLruCache(10240)
	}
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || Startwiths(ip, "127") || Startwiths(ip, "192") || Startwiths(ip, "172") || Startwiths(ip, "10.") {
		return "内网"
	}

	addr, find := _ipCache.Get(ip)
	if find {
		return addr
	}

	// {"ip":"x.x.x.x","pro":"浙江省","proCode":"330000","city":"杭州市","cityCode":"330100","region":"","regionCode":"0","addr":"浙江省杭州市 电信","regionNames":"","err":""}
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bt, err := HttpGetJsonTimeout(url, 3*time.Second)
	if err != nil {
		return ""
	}

	d := &ipCityResult{}
	err = json.Unmarshal(GbkToUtf8(bt), d)
	if err != nil {
		return ""
	}

	if d.Addr != "" {
		_ipCache.Add(ip, d.Addr)
		return d.Addr
	} else if d.Pro != "" && d.Pro == d.City {
		_ipCache.Add(ip, d.Pro)
		return d.Pro
	} else {
		_ipCache.Add(ip, Trim(d.Pro+d.City))
		return Trim(d.Pro + d.City)
	}
}

// 获取ip地址信息不含ip
func GetCityByIp_ip9(ip string) string {
	if _ipCache == nil {
		_ipCache = NewLruCache(10240)
	}
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || Startwiths(ip, "127") || Startwiths(ip, "192") || Startwiths(ip, "172") || Startwiths(ip, "10.") {
		return "内网"
	}

	addr, find := _ipCache.Get(ip)
	if find {
		return addr
	}

	// {"ret":200,"data":{"ip":"xxx.xxx.xxx.xxx","country":"中国","country_code":"cn","prov":"广东","city":"深圳","city_code":"shenzhen","city_short_code":"sz","area":"龙岗","post_code":"518116","area_code":"0755","isp":"中国电信","lng":"114.24771","lat":"22.71986","long_ip":2032275147,"big_area":"华南"},"qt":0}
	url := "https://ip9.com.cn/get?ip=" + ip
	bt, err := HttpGetJsonTimeout(url, 3*time.Second)
	if err != nil {
		return ""
	}

	d := &ipInfoResponse{}
	err = json.Unmarshal(bt, d)
	if err != nil {
		return ""
	}

	rsAddr := getIpStr(d)

	if rsAddr != "" {
		_ipCache.Add(ip, rsAddr)
	}
	return rsAddr
}

func getIpStr(d *ipInfoResponse) string {
	var rs = ""
	if d.Ret != 200 {
		return ""
	}

	rs += d.Data.Country
	rs += d.Data.Province
	rs += d.Data.City
	rs += d.Data.Area

	if d.Data.ISP != "" {
		rs += " " + d.Data.ISP
	}

	rs = ReplaceAll(rs, "中国", "中国")
	rs = ReplaceAll(rs, "北京北京", "北京")
	rs = ReplaceAll(rs, "天津天津", "天津")
	rs = ReplaceAll(rs, "上海上海", "上海")
	rs = ReplaceAll(rs, "重庆重庆", "重庆")
	return rs
}

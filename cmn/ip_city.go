package cmn

import "encoding/json"

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
	if _ipCache == nil {
		_ipCache = NewLruCache(128)
	}
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" || Startwiths(ip, "192") || Startwiths(ip, "172") || Startwiths(ip, "10.") {
		return "内网"
	}

	addr, find := _ipCache.Get(ip)
	if find {
		return addr
	}

	// {"ip":"x.x.x.x","pro":"浙江省","proCode":"330000","city":"杭州市","cityCode":"330100","region":"","regionCode":"0","addr":"浙江省杭州市 电信","regionNames":"","err":""}
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bt, err := HttpGetJson(url)
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

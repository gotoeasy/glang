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

// 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" || Startwiths(ip, "192") || Startwiths(ip, "172") || Startwiths(ip, "10.") {
		return "内网"
	}

	// {"ip":"x.x.x.x","pro":"浙江省","proCode":"330000","city":"杭州市","cityCode":"330100","region":"","regionCode":"0","addr":"浙江省杭州市 电信","regionNames":"","err":""}
	url := "https://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bt, err := HttpGetJson(url)
	if err != nil {
		Error(err)
		return ""
	}

	d := &ipCityResult{}
	err = json.Unmarshal(GbkToUtf8(bt), d)
	if err != nil {
		Error(err)
		return ""
	}

	if d.Pro != "" && d.Pro == d.City {
		return d.Pro
	}

	return Trim(d.Pro + d.City)
}

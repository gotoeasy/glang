package cmn

import "net"

var localIpAddres string
var localHostName string

// 取本机IP地址（IPv4）
func GetLocalIp() string {
	if localIpAddres != "" {
		return localIpAddres
	}

	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				localIpAddres = ipnet.IP.String()
			}
		}
	}
	return localIpAddres
}

// 取本机名
func GetLocalHostName() string {
	if localHostName != "" {
		return localHostName
	}

	info, err := MeasureHost()
	if err != nil {
		localHostName = info.Hostname
	}
	return localHostName
}

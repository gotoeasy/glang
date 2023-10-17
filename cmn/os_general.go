package cmn

import (
	"net"
	"os/exec"
	"strings"
)

var localIpAddres string
var localInternalIpAddres string
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
				if Startwiths(localIpAddres, "172.") || Startwiths(localIpAddres, "192.") || Startwiths(localIpAddres, "10.") {
					break // 内部IP优先
				}
			}
		}
	}
	return localIpAddres
}

// 取本机IP地址（IPv4），eth0优先
func GetPreferredLocalIPv4() string {
	if localInternalIpAddres != "" {
		return localInternalIpAddres // 用缓存
	}
	interfaces, err := net.Interfaces()
	if err != nil {
		localInternalIpAddres = GetLocalIp()
		return localInternalIpAddres
	}

	// 优先选择eth0接口的IPv4地址
	for _, iface := range interfaces {
		if strings.HasPrefix(iface.Name, "eth0") {
			addrs, err := iface.Addrs()
			if err != nil {
				localInternalIpAddres = GetLocalIp()
				return localInternalIpAddres
			}

			for _, addr := range addrs {
				ipnet, ok := addr.(*net.IPNet)
				if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					rs := ipnet.IP.String()
					if Startwiths(rs, "172.") || Startwiths(rs, "192.") || Startwiths(rs, "10.") {
						localInternalIpAddres = rs
						return localInternalIpAddres // 理想情况下是从这里取到
					}
				}
			}
		}
	}

	localInternalIpAddres = GetLocalIp()
	return localInternalIpAddres
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

	if localHostName == "" {
		localHostName = getHostname()
	}
	return localHostName
}

func getHostname() string {
	cmd := exec.Command("hostname")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	hostname := Trim(string(output))
	return hostname
}

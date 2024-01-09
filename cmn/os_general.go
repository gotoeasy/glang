package cmn

import (
	"net"
	"os/exec"
	"strings"
)

var localIpAddres string
var localHostName string

// 按顺序取本机IP地址（IPv4），优先度 eth0 ip > 192.* > 172.* > 10.* > 其他
func GetLocalIp() string {
	return getPreferredLocalIPv4()
}

// 取本机IP地址（IPv4），优先度 192.* > 172.* > 10.* > 其他
func getLocalIpAddres() string {
	if localIpAddres != "" {
		return localIpAddres
	}

	addrs, err := net.InterfaceAddrs()
	if err == nil {
		ip192, ip172, ip10, ipOther := "", "", "", ""
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				if Startwiths(ip, "192.") {
					ip192 = ip
					break // 192.*的IP优先
				}
				if Startwiths(ip, "172.") {
					ip172 = ip
				} else if Startwiths(ip, "10.") {
					ip10 = ip
				} else if ipOther != "" {
					ipOther = ip
				}
			}
		}

		// 192.* > 172.* > 10.* > 其他
		if ipOther != "" {
			localIpAddres = ipOther
		}
		if ip10 != "" {
			localIpAddres = ip10
		}
		if ip172 != "" {
			localIpAddres = ip172
		}
		if ip192 != "" {
			localIpAddres = ip192
		}
	}
	return localIpAddres
}

// 取本机IP地址（IPv4），优先度 eth0 ip > 192.* > 172.* > 10.* > 其他
func getPreferredLocalIPv4() string {
	if localIpAddres != "" {
		return localIpAddres // 用缓存
	}
	interfaces, err := net.Interfaces()
	if err != nil {
		localIpAddres = getLocalIpAddres()
		return localIpAddres
	}

	// 优先选择eth0接口的IPv4地址
	for _, iface := range interfaces {
		if strings.HasPrefix(iface.Name, "eth0") {
			addrs, err := iface.Addrs()
			if err != nil {
				localIpAddres = getLocalIpAddres()
				return localIpAddres
			}

			for _, addr := range addrs {
				ipnet, ok := addr.(*net.IPNet)
				if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					rs := ipnet.IP.String()
					if Startwiths(rs, "172.") || Startwiths(rs, "192.") || Startwiths(rs, "10.") {
						localIpAddres = rs
						return localIpAddres // 理想情况下是从这里取到
					}
				}
			}
		}
	}

	localIpAddres = getLocalIpAddres()
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

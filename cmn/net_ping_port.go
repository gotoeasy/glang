package cmn

import (
	"net"
	"time"
)

// 检测本机指定端口是否打开中
func IsPortOpening(port string) bool {
	return IsServerPortOpening("127.0.0.1", port)
}

// 检测指定服务指定端口是否打开中
func IsServerPortOpening(ip string, port string) bool {
	conn, err := net.DialTimeout("tcp", ip+":"+port, time.Millisecond*500)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

package cmn

import (
	"os"
	"path/filepath"

	"github.com/gotoeasy/glang/cmn"
)

// 【win】在 C:/Windows/System32/drivers/etc/hosts 中添加指定映射
func HostsAddDomain(ip string, domain string) error {
	hostsfile := "C:\\Windows\\System32\\drivers\\etc\\hosts"
	txt, _ := cmn.ReadFileString(hostsfile)
	lines := cmn.Split(txt, "\n")
	flg := false
	for i := 0; i < len(lines); i++ {
		if (ip + " " + domain) == cmn.ReplaceAllSpace(cmn.Trim(lines[i]), " ") {
			flg = true
			break
		}
	}
	if flg {
		return nil
	}

	// 首行添加映射
	return cmn.WriteFileString(hostsfile, ip+"    "+domain+"\r\n"+txt)
}

func GetUserHomeDir() string {
	dir, _ := os.UserHomeDir()
	return dir
}

// 【win】Window系统Startup目录
func GetUserStartupDir() string {
	return filepath.Join(GetUserHomeDir(), `AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`)
}

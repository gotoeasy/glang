package cmn

// 在 C:/Windows/System32/drivers/etc/hosts 中添加指定映射
func HostsAddDomain(ip string, domain string) error {
	hostsfile := "C:\\Windows\\System32\\drivers\\etc\\hosts"
	txt, _ := ReadFileString(hostsfile)
	lines := Split(txt, "\n")
	flg := false
	for i := 0; i < len(lines); i++ {
		if (ip + " " + domain) == ReplaceAllSpace(Trim(lines[i]), " ") {
			flg = true
			break
		}
	}
	if flg {
		return nil
	}

	// 首行添加映射
	return WriteFileString(hostsfile, ip+"    "+domain+"\r\n"+txt)
}

package cmn

import "strings"

// 判断字符串是否在切片中
func IncludesStr(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// 判断字符串是否在切片中（忽略大小写）
func IncludesStrIgnoreCase(slice []string, str string) bool {
	for _, s := range slice {
		if strings.EqualFold(str, s) {
			return true
		}
	}
	return false
}

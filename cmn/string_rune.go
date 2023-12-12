package cmn

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// 按文字计算字符串长度
func Len(str string) int {
	return len([]rune(str))
}

// 取左文字
func Left(str string, length int) string {
	srune := []rune(str)
	lenr := len(srune)
	if lenr <= length {
		return str
	}

	var rs string
	for i := 0; i < length; i++ {
		rs += string(srune[i])
	}
	return rs
}

// 取右文字
func Right(str string, length int) string {
	srune := []rune(str)
	lenr := len(srune)
	if lenr <= length {
		return str
	}

	var rs string
	for i := lenr - length; i < lenr; i++ {
		rs += string(srune[i])
	}
	return rs
}

// 去除两边空格
func Trim(str string) string {
	return strings.TrimSpace(str)
}

// 去除左前缀
func TrimPrefix(str string, prefix string) string {
	return strings.TrimPrefix(str, prefix)
}

// 判断是否空白
func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

// 判断是否指定前缀
func Startwiths(str string, startstr string, ignoreCase ...bool) bool {
	lstr := Left(str, Len(startstr))
	if len(ignoreCase) > 0 && ignoreCase[0] {
		return EqualsIngoreCase(lstr, startstr)
	}
	return lstr == startstr
}

// 判断是否指定后缀
func Endwiths(str string, endstr string, ignoreCase ...bool) bool {
	rstr := Right(str, Len(endstr))
	if len(ignoreCase) > 0 && ignoreCase[0] {
		return EqualsIngoreCase(rstr, endstr)
	}
	return rstr == endstr
}

// 按文字截取字符串
func SubString(str string, start int, end int) string {
	srune := []rune(str)
	slen := len(srune)
	if start >= slen || start >= end || start < 0 {
		return ""
	}

	rs := ""
	for i := start; i < slen && i < end; i++ {
		rs += string(srune[i])
	}
	return rs
}

// 查找文字下标
func IndexOf(str string, substr string) int {
	idx := strings.Index(str, substr)
	if idx < 0 {
		return idx
	}
	return utf8.RuneCountInString(str[:idx])
}

// 判断是否包含（区分大小写）
func Contains(str string, substr string) bool {
	return IndexOf(str, substr) >= 0
}

// 判断是否包含（忽略大小写）
func ContainsIngoreCase(str string, substr string) bool {
	return IndexOf(ToLower(str), ToLower(substr)) >= 0
}

// 判断是否相同（忽略大小写）
func EqualsIngoreCase(str1 string, str2 string) bool {
	return ToLower(str1) == ToLower(str2)
}

// 转小写
func ToLower(str string) string {
	return strings.ToLower(str)
}

// 转大写
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// 重复
func Repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

// 左补足
func PadLeft(str string, pad string, length int) string {
	if length < Len(str) {
		return str
	}
	s := Repeat(pad, length) + str
	max := Len(s)
	return SubString(s, max-length, max)
}

// 右补足
func PadRight(str string, pad string, length int) string {
	if length < Len(str) {
		return str
	}
	s := str + Repeat(pad, length)
	return SubString(s, 0, length)
}

// 替换
func Replace(str string, old string, new string, n int) string {
	return strings.Replace(str, old, new, n)
}

// 全部替换
func ReplaceAll(str string, old string, new string) string {
	return strings.ReplaceAll(str, old, new)
}

// 全部替换连续的空白
func ReplaceAllSpace(str string, new string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(str, new)
}

// 反转
func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 字符串切割
func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

// 字符串数组拼接为字符串
func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

// 字符串去重
func Unique(strs []string) []string {
	m := make(map[string]struct{}, 0)
	newS := make([]string, 0)
	for _, i2 := range strs {
		if _, ok := m[i2]; !ok {
			newS = append(newS, i2)
			m[i2] = struct{}{}
		}
	}
	return newS
}

// 首字母转大写
func Titlelize(str string) string {
	return ToUpper(Left(str, 1)) + Right(str, Len(str)-1)
}

// 驼峰转全小写下划线(已含下划线时直接转小写)
func CamelToUnderline(str string) string {
	if Contains(str, "_") {
		return ToLower(str)
	}

	var rs []rune
	for i, r := range str {
		if i == 0 {
			rs = append(rs, r)
		} else {
			if unicode.IsUpper(r) {
				rs = append(rs, '_')
			}
			rs = append(rs, r)
		}
	}
	return ToLower(string(rs))
}

// 下划线转驼峰(无下划线时不转换)
func UnderlineToCamel(str string) string {
	if !Contains(str, "_") {
		return str
	}

	ary := Split(ToLower(str), "_")
	var rs string
	for i := 0; i < len(ary); i++ {
		rs += Titlelize(ary[i])
	}
	return rs
}

// 按K或M或G单位显示，保留1位小数
func GetSizeInfo(size uint64) string {
	if size > 1024*1024*1024 {
		return fmt.Sprintf("%.1fG", float64(size)/1024/1024/1024)
	}
	if size > 1024*1024 {
		return fmt.Sprintf("%.1fM", float64(size)/1024/1024)
	}
	return fmt.Sprintf("%.1fK", float64(size)/1024)
}

// 按容易理解的单位表示时间
func GetTimeInfo(milliseconds int64) string {
	seconds := milliseconds / 1000
	minutes := seconds / 60
	hours := minutes / 60

	if hours > 0 {
		minutes %= 60
		return fmt.Sprintf("%d小时%d分", hours, minutes)
	} else if minutes > 0 {
		seconds %= 60
		return fmt.Sprintf("%d分%d秒", minutes, seconds)
	} else if seconds > 0 {
		return fmt.Sprintf("%d秒", seconds)
	}

	return fmt.Sprintf("%d毫秒", milliseconds)
}

package cmn

import (
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func Len(str string) int {
	return utf8.RuneCountInString(str)
}

func Left(str string, length int) string {
	if Len(str) <= length {
		return str
	}

	var rs string
	for i, s := range str {
		if i == length {
			break
		}
		rs = rs + string(s)
	}
	return rs
}

func Right(str string, length int) string {
	lenr := Len(str)
	if lenr <= length {
		return str
	}

	var rs string
	start := lenr - length
	for i, s := range str {
		if i >= start {
			rs = rs + string(s)
		}
	}
	return rs
}

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func Startwiths(str string, startstr string) bool {
	lstr := Left(str, Len(startstr))
	return lstr == startstr
}

func Endwiths(str string, endstr string) bool {
	rstr := Right(str, Len(endstr))
	return rstr == endstr
}

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

func IndexOf(str string, substr string) int {
	idx := strings.Index(str, substr)
	return utf8.RuneCountInString(str[:idx])
}

func Contains(str string, substr string) bool {
	return IndexOf(str, substr) >= 0
}

func ContainsIngoreCase(str string, substr string) bool {
	return IndexOf(ToLower(str), ToLower(substr)) >= 0
}

func EqualsIngoreCase(str1 string, str2 string) bool {
	return ToLower(str1) == ToLower(str2)
}

func ToLower(str string) string {
	return strings.ToLower(str)
}

func ToUpper(str string) string {
	return strings.ToUpper(str)
}

func Repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

func PadLeft(str string, pad string, length int) string {
	if length < Len(str) {
		return str
	}
	s := Repeat(pad, length) + str
	max := Len(s)
	return SubString(s, max-length, max)
}

func PadRight(str string, pad string, length int) string {
	if length < Len(str) {
		return str
	}
	s := str + Repeat(pad, length)
	return SubString(s, 0, length)
}

func Replace(str string, old string, new string, n int) string {
	return strings.Replace(str, old, new, n)
}

func ReplaceAll(str string, old string, new string) string {
	return strings.ReplaceAll(str, old, new)
}

func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

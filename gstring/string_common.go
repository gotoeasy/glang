package gstring

import (
	"strings"
	"unicode/utf8"
)

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

func Startwiths(str string, startstr string) bool {

	if startstr == "" || str == startstr {
		return true
	}

	strs := []rune(str)
	tmps := []rune(startstr)
	lens := len(strs)
	lentmp := len([]rune(tmps))
	if lens < lentmp {
		return false
	}

	for i := 0; i < lentmp; i++ {
		if tmps[i] != strs[i] {
			return false
		}
	}

	return true
}

func Endwiths(str string, endstr string) bool {

	if endstr == "" || str == endstr {
		return true
	}

	strs := []rune(str)
	ends := []rune(endstr)
	lens := len(strs)
	lene := len(ends)
	if lens < lene {
		return false
	}

	dif := lens - lene
	for i := lene - 1; i >= 0; i-- {
		if strs[dif+i] != ends[i] {
			return false
		}
	}

	return true
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

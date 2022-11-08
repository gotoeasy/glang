package cmn

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 随机数
func RandomInt(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// 随机数
func RandomUint32() uint32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
}

// 绝对值
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 绝对值
func AbsInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

// 四舍五入保留指定位数(0-16)的小数
func Round(num float64, digit int) float64 {
	if digit < 0 {
		digit = 0
	}
	if digit > 16 {
		digit = 16
	}
	rs, _ := strconv.ParseFloat(fmt.Sprintf("%."+IntToString(digit)+"f", num), 64)
	return rs
}

// 四舍五入保留1位小数
func Round1(num float64) float64 {
	rs, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", num), 64)
	return rs
}

// 四舍五入保留2位小数
func Round2(num float64) float64 {
	rs, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return rs
}

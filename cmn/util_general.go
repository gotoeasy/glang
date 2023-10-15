package cmn

import (
	"time"
)

// 计算函数执行时间（毫秒）
func ExecTime(fn func()) int64 {
	start := time.Now()
	fn()
	return time.Since(start).Milliseconds()
}

// 执行回调函数，错误时重试
func Retry(callback func() error, retryTimes int, duration time.Duration) (err error) {
	for i := 1; i <= retryTimes; i++ {
		if err = callback(); err != nil {
			time.Sleep(duration)
			continue
		}
		return
	}
	return
}

// 版本号格式转换便于比较大小，格式不符时返回原版本，例 v1.2.3 => v01.002.003
func NormalizeVer(ver string) string {
	ary1 := Split(ver, ".")
	if len(ary1) != 3 || !IsDigit(ary1[1]) || !IsDigit(ary1[2]) {
		return ver
	}
	mainVer := ReplaceAll(ary1[0], "v", "")
	if IsDigit(mainVer) {
		ary1[0] = "v" + Right(IntToString(100+StringToInt(mainVer, 0)), 2)
	}
	ary1[1] = Right(IntToString(1000+StringToInt(ary1[1], 0)), 3)
	ary1[2] = Right(IntToString(1000+StringToInt(ary1[2], 0)), 3)
	return Join(ary1, ".")
}

// 条件参数真时返回字符串1，否则返回字符串2
func IifStr(condtion bool, s1 string, s2 string) string {
	if condtion {
		return s1
	}
	return s2
}

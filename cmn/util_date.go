package cmn

import (
	"time"
)

// 当日的yyyymmdd格式
func Today() string {
	return time.Now().Format("20060102")
}

// 当前日期加减天数后的yyyymmdd格式
func GetYyyymmdd(addDays int) string {
	return time.Now().AddDate(0, 0, addDays).Format("20060102")
}

// 系统时间的yyyymmdd HH:MM:SS格式
func GetYyyymmddHHMMSS() string {
	return time.Now().Format("20060102 15:04:05")
}

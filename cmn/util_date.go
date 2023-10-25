package cmn

import (
	"time"
)

// 日期格式，使用常量 FMT_XXX
type DateFormat string

// 日期格式模板
const (
	FMT_MM_DD                   = "MM-dd"
	FMT_YYYYMM                  = "yyyyMM"
	FMT_YYYY_MM                 = "yyyy-MM"
	FMT_YYYY_MM_DD              = "yyyy-MM-dd"
	FMT_YYYYMMDD                = "yyyyMMdd"
	FMT_YYYYMMDDHHMMSS          = "yyyyMMddHHmmss"
	FMT_YYYYMMDDHHMM            = "yyyyMMddHHmm"
	FMT_YYYYMMDDHH              = "yyyyMMddHH"
	FMT_YYMMDDHHMM              = "yyMMddHHmm"
	FMT_MM_DD_HH_MM             = "MM-dd HH:mm"
	FMT_MM_DD_HH_MM_SS          = "MM-dd HH:mm:ss"
	FMT_YYYY_MM_DD_HH_MM        = "yyyy-MM-dd HH:mm"
	FMT_YYYY_MM_DD_HH_MM_SS     = "yyyy-MM-dd HH:mm:ss"
	FMT_YYYY_MM_DD_HH_MM_SS_SSS = "yyyy-MM-dd HH:mm:ss.SSS"

	FMT_MM_DD_EN                   = "MM/dd"
	FMT_YYYY_MM_EN                 = "yyyy/MM"
	FMT_YYYY_MM_DD_EN              = "yyyy/MM/dd"
	FMT_MM_DD_HH_MM_EN             = "MM/dd HH:mm"
	FMT_MM_DD_HH_MM_SS_EN          = "MM/dd HH:mm:ss"
	FMT_YYYY_MM_DD_HH_MM_EN        = "yyyy/MM/dd HH:mm"
	FMT_YYYY_MM_DD_HH_MM_SS_EN     = "yyyy/MM/dd HH:mm:ss"
	FMT_YYYY_MM_DD_HH_MM_SS_SSS_EN = "yyyy/MM/dd HH:mm:ss.SSS"

	FMT_MM_DD_CN               = "MM月dd日"
	FMT_YYYY_MM_CN             = "yyyy年MM月"
	FMT_YYYY_MM_DD_CN          = "yyyy年MM月dd日"
	FMT_MM_DD_HH_MM_CN         = "MM月dd日 HH:mm"
	FMT_MM_DD_HH_MM_SS_CN      = "MM月dd日 HH:mm:ss"
	FMT_YYYY_MM_DD_HH_MM_CN    = "yyyy年MM月dd日 HH:mm"
	FMT_YYYY_MM_DD_HH_MM_SS_CN = "yyyy年MM月dd日 HH:mm:ss"

	FMT_HH_MM       = "HH:mm"
	FMT_HH_MM_SS    = "HH:mm:ss"
	FMT_HH_MM_SS_MS = "HH:mm:ss.SSS"
)

// 当日的yyyymmdd格式
func Today() string {
	return time.Now().Format("20060102")
}

// 当前日期加减天数后的yyyymmdd格式
func GetYyyymmdd(addDays int) string {
	return time.Now().AddDate(0, 0, addDays).Format("20060102")
}

// 返回年月周yyyymm+week
func GetYearMonthWeek() string {
	currentTime := time.Now()
	year := currentTime.Format("2006")
	month := currentTime.Format("01")
	_, week := currentTime.ISOWeek()        // 系统日期属于当月第几周
	return year + month + IntToString(week) // 拼接(年、月、周)返回
}

// 格式化日期
func FormatDate(date time.Time, fmt DateFormat) string {
	layout := string(fmt)
	layout = Replace(layout, "yyyy", "2006", 1)
	layout = Replace(layout, "yy", "06", 1)
	layout = Replace(layout, "MM", "01", 1)
	layout = Replace(layout, "dd", "02", 1)
	layout = Replace(layout, "HH", "15", 1)
	layout = Replace(layout, "mm", "04", 1)
	layout = Replace(layout, "ss", "05", 1)
	layout = Replace(layout, "SSS", "000", 1)

	return date.Format(layout)
}

// 格式化系统日期
func FormatSystemDate(fmt DateFormat) string {
	return FormatDate(time.Now(), fmt)
}

// 指定格式的字符串转日期
func ParseDate(date string, fmt DateFormat) (time.Time, error) {
	return time.Parse(string(fmt), date)
}

// 系统时间 yyyy-MM-dd HH:mm:ss.SSS
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05.000")
}

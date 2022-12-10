package cmn

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// float64 转 string，保留指定位数小数，后续小数位舍去
//
// Float64ToStringRoundDown(1234567890123456,2) -> 1234567890123456.00
// Float64ToStringRoundDown(999999999999.999,2) -> 999999999999.99
func Float64ToStringRoundDown(num float64, digit int32) string {
	return decimal.NewFromFloat(num).RoundDown(digit).StringFixedBank(digit)
}

// float64 转 string (注：千万亿大数时可能进位，不绝对)
func Float64ToString(num float64) string {
	return decimal.NewFromFloat(num).String()
}

// float64 转 string，四舍五入保留或补足指定位数小数
func FormatRound(num float64, digit int32) string {
	return decimal.NewFromFloat(num).StringFixed(digit)
}

// float64 转 string，整数部分3位一撇，四舍五入保留或补足指定位数小数
func FormatAmountRound(num float64, digit int32) string {
	s := FormatRound(num, digit)
	ary := Split(s, ".")

	// 整数部分三位一撇
	a := Split(ary[0], "")
	cnt := 0
	rs := ""
	for i := len(a) - 1; i >= 0; i-- {
		if cnt == 3 {
			rs = "," + rs
			cnt = 0
		}
		rs = a[i] + rs
		cnt++
	}

	// 有小数时拼接小数部分
	if len(ary) == 2 {
		rs = rs + "." + ary[1]
	}
	return Replace(rs, "-,", "-", 1) // 特殊情况处理（如 -,123,456.789 => -123,456.789）
}

// 四舍五入保留指定位数(0-16)的小数
func Round(num float64, digit int32) float64 {
	if digit < 0 {
		digit = 0
	}
	if digit > 16 {
		digit = 16
	}
	rs, _ := decimal.NewFromFloat(num).Round(digit).Float64()
	return rs
}

// 四舍五入保留1位小数
func Round1(num float64) float64 {
	rs, _ := decimal.NewFromFloat(num).Round(1).Float64()
	return rs
}

// 四舍五入保留2位小数
func Round2(num float64) float64 {
	rs, _ := decimal.NewFromFloat(num).Round(2).Float64()
	return rs
}

// 保留指定位数(0-16)的小数（后面小数舍去）
func RoundDown(num float64, digit int32) float64 {
	if digit < 0 {
		digit = 0
	}
	if digit > 16 {
		digit = 16
	}
	rs, _ := decimal.NewFromFloat(num).RoundDown(digit).Float64()
	return rs
}

// 金额数字转人民币大写（百万亿级别正常，超过范围可能转数字字符串出现进位）
func Float64ToCny(val float64) string {
	return AmountToCny(Float64ToStringRoundDown(val, 2))
}

// 金额数字转人民币大写（最大支持千万亿，小数只精确到分，分以下舍去。超过支持的最大值时原样返回不转换）
//
// 1234567890123456.789  -> 壹仟贰佰叁拾肆万伍仟陆佰柒拾捌亿玖仟零壹拾贰万叁仟肆佰伍拾陆元柒角捌分
// -1234567890123456.78  -> 负壹仟贰佰叁拾肆万伍仟陆佰柒拾捌亿玖仟零壹拾贰万叁仟肆佰伍拾陆元柒角捌分
// 12345.7               -> 壹万贰仟叁佰肆拾伍元柒角整
// 12345.0               -> 壹万贰仟叁佰肆拾伍元整
// 12345                 -> 壹万贰仟叁佰肆拾伍元整
// 9002300040            -> 玖拾亿零贰佰叁拾万零肆拾元整
// 9002300043.0          -> 玖拾亿零贰佰叁拾万零肆拾叁元整
// 12345678901234567.781 -> 12345678901234567.781
func AmountToCny(val string) string {

	ary := Split(val, ".")
	num := ary[0]
	if len(ary) == 1 {
		num += "00"
	} else if len(ary) == 2 {
		if len(ary[1]) > 2 {
			num += Left(ary[1], 2)
		} else {
			num += PadRight(ary[1], "0", 2)
		}
	}

	pre := ""
	if Startwiths(num, "-") {
		pre = "负"
		num = ReplaceAll(num, "-", "")
	}

	if len(num) > 18 {
		return val
	}

	chineseMap := []string{"分", "角", "元", "拾", "佰", "仟", "万", "拾", "佰", "仟", "亿", "拾", "佰", "仟", "万", "拾", "佰", "仟"}
	chineseNum := []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	listNum := []int{}

	for _, s := range strings.Split(Reverse(num), "") {
		listNum = append(listNum, StringToInt(s, 0))
	}
	n := len(listNum)
	chinese := ""
	for i := n - 1; i >= 0; i-- {
		chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
	}

	for {
		copychinese := chinese
		copychinese = strings.Replace(copychinese, "零万", "万", 1)
		copychinese = strings.Replace(copychinese, "零亿", "亿", 1)
		copychinese = strings.Replace(copychinese, "零拾", "零", 1)
		copychinese = strings.Replace(copychinese, "零佰", "零", 1)
		copychinese = strings.Replace(copychinese, "零仟", "零", 1)
		copychinese = strings.Replace(copychinese, "零角", "零", 1)
		copychinese = strings.Replace(copychinese, "零零", "零", 1)
		copychinese = strings.Replace(copychinese, "零元", "元", 1)
		copychinese = strings.Replace(copychinese, "零分", "", 1)

		if copychinese == chinese {
			break
		} else {
			chinese = copychinese
		}
	}

	rs := pre + chinese
	if !Endwiths(rs, "分") {
		rs += "整"
	}
	return rs
}

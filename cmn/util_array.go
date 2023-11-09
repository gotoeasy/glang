package cmn

// 判断数组中是否包含指定字符串
func ContainsItem(ary []string, str string, ingnoreCase ...bool) bool {
	ingnore := false
	if len(ingnoreCase) > 0 {
		ingnore = ingnoreCase[0]
	}
	for i, max := 0, len(ary); i < max; i++ {
		if ingnore {
			if EqualsIngoreCase(ary[i], str) {
				return true
			}
		} else {
			if ary[i] == str {
				return true
			}
		}
	}
	return false
}

package cmn

import (
	"testing"
)

func Test_sego(t *testing.T) {

	seg := NewTokenizerSego("")
	ws := seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造，Java和Go都学得不错，Java和Go都不错")
	Info(ws)
}

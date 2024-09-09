package cmn

import (
	"testing"
)

func Test_sego(t *testing.T) {
	seg := NewTokenizerSego() // 没有指定字典，将会简单的单字切分
	ws := seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造，Java和Go都学得不错，Java和Go都不错的")
	Info(ws)

	seg.IngoreWords("的", "于")
	ws = seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造，Java和Go都学得不错，Java和Go都不错")
	Info(ws)

}

func Test_GetHtmlText(t *testing.T) {
	s := GetHtmlText(`<ds sas="ddddd<>ddd"/>dsads111<div>ssss222</div>dsasss<d/ssf(9sf)>333ssd &nbsp`)
	Info(s)
}

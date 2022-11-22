package cmn

import (
	"testing"
)

func Test_sego(t *testing.T) {

	seg := NewTokenizerSego("")
	ws := seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造，Java和Go都学得不错，Java和Go都不错的")
	Info(ws)

	seg.IngoreWords("的", "于")
	ws = seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造，Java和Go都学得不错，Java和Go都不错")
	Info(ws)

}

func Test_jiebago(t *testing.T) {
	jieba := NewTokenizerJiebago("")
	ws := jieba.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造，Java和Go都学得不错，Java和Go都不错的")
	Info(ws)

	jieba.IngoreWords("的", "于")
	ws = jieba.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造，Java和Go都学得不错，Java和Go都不错的")
	Info(ws)

}

func Test_GetHtmlText(t *testing.T) {
	s := GetHtmlText(`<ds sas="ddddd<>ddd"/>dsads111<div>ssss222</div>dsasss<d/ssf(9sf)>333ssd &nbsp`)
	Info(s)
}

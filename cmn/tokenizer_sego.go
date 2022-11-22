package cmn

import (
	"sync"

	"github.com/huichen/sego"
)

type TokenizerSego struct {
	segmenter      sego.Segmenter
	mapIngoreWords map[string]bool
}

var _segmenterSego *TokenizerSego
var _segmenterSegoMu sync.Mutex

// 创建中文分词器（sego）
// 参数dicFile为字典文件，传入空时默认为"data/dictionary.txt"
func NewTokenizerSego(dicFile string) *TokenizerSego {
	if _segmenterSego != nil {
		return _segmenterSego
	}
	_segmenterSegoMu.Lock()
	defer _segmenterSegoMu.Unlock()
	if _segmenterSego != nil {
		return _segmenterSego
	}

	// 载入词典
	if IsBlank(dicFile) {
		dicFile = "data/dictionary.txt"
	}
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(dicFile)

	_segmenterSego = &TokenizerSego{
		segmenter:      segmenter,
		mapIngoreWords: make(map[string]bool),
	}
	// 初始化默认忽略的字符
	ingoreChars := "`~!@# $%^&*()-_=+[{]}\\|;:'\",<.>/?，。《》；：‘　’“”、|】｝【｛＋－—（）×＆…％￥＃＠！～·\t\r\n"
	for _, s := range ingoreChars {
		_segmenterSego.mapIngoreWords[string(s)] = true
	}

	return _segmenterSego
}

// 设定忽略词（比如分词结果不想包含无效词“的”或一些敏感词时，可以这里设定）
func (t *TokenizerSego) IngoreWords(str ...string) {
	for _, s := range str {
		t.mapIngoreWords[s] = true
	}
}

// 按搜索引擎模式进行分词（自动去重、去标点符号、忽略大小写）
func (t *TokenizerSego) CutForSearch(str string) []string {
	segs := t.segmenter.Segment(StringToBytes(ToLower(str))) // 转小写后再分词
	ws := sego.SegmentsToSlice(segs, true)

	var rs []string
	var mapStr = make(map[string]string)
	for _, w := range ws {
		if _, has := t.mapIngoreWords[w]; has {
			continue // 去忽略词
		}
		if _, has := mapStr[w]; has {
			continue // 去重
		}
		mapStr[w] = ""
		rs = append(rs, w)
	}
	return rs
}

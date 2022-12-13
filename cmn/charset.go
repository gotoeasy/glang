package cmn

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// gbk转utf8
func GbkToUtf8(gbk []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(gbk), simplifiedchinese.GBK.NewDecoder())
	d, err := io.ReadAll(reader)
	if err != nil {
		Error(err)
		return gbk
	}

	return d
}

// utf8转gbk
func Utf8ToGbk(utf8 []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(utf8), simplifiedchinese.GBK.NewEncoder())
	d, err := io.ReadAll(reader)
	if err != nil {
		Error(err)
		return utf8
	}

	return d
}

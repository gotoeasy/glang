package img

import (
	"bytes"
	"image/jpeg"

	"github.com/gotoeasy/glang/cmn"
	"golang.org/x/image/bmp"
)

// bmp文件转jpg文件
func ImgBmpToJpg(buf []byte, o *jpeg.Options) []byte {

	img, err := bmp.Decode(bytes.NewReader(buf))
	if err != nil {
		cmn.Error(err)
		return buf
	}

	if o == nil {
		o = &jpeg.Options{Quality: 80}
	}

	newBuf := bytes.Buffer{}
	err = jpeg.Encode(&newBuf, img, o)
	if err != nil {
		cmn.Error(err)
		return buf
	}

	return newBuf.Bytes()
}

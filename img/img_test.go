package img

import (
	"testing"

	"github.com/gotoeasy/glang/cmn"
)

func Test_img(t *testing.T) {

	ImgResize("d:\\src.png", "d:\\dist1.png", 1024, 1024, nil)

	cmn.Info(ImgBlur("d:\\src.png", "d:\\dist2.png", 10))

}

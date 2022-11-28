package cmn

import (
	"testing"
)

func Test_color(t *testing.T) {
	Info(HexToRgb("#0081cc"))
	Info(HexToRgb("#aabbcc"))
	Info(RgbToHex(0, 129, 204))
	Info(RgbToHex(170, 187, 204))
}

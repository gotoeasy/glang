package cmn

import (
	"testing"
)

func Test_uid(t *testing.T) {
	for i := 0; i < 10; i++ {
		Info(ULID())
	}
}

func Test_decimal(t *testing.T) {
	Info(FormatAmountRound(1, 3))
	Info(FormatAmountRound(1, 0))
	Info(FormatAmountRound(1, -1))
	Info(FormatAmountRound(-121345678.1, 3))
}

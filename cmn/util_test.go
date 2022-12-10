package cmn

import (
	"testing"
)

func Test_uid(t *testing.T) {
	for i := 0; i < 10; i++ {
		Info(ULID())
	}
}

func Test_str(t *testing.T) {
	Info(Left("啊啊啊哦哦哦呃呃呃", 3))
	Info(Right("啊啊啊哦哦哦呃呃呃", 3))
}

func Test_decimal(t *testing.T) {
	Info(FormatAmountRound(1, 3))
	Info(FormatAmountRound(1, 0))
	Info(FormatAmountRound(1, -1))
	Info(FormatAmountRound(-121345678.1, 3))

	Info(AmountToCny("1234567890123456.789"))
	Info(AmountToCny("-1234567890123456.78"))
	Info(AmountToCny("12345.7"))
	Info(AmountToCny("12345.0"))
	Info(AmountToCny("12345"))
	Info(AmountToCny("9002300040"))
	Info(AmountToCny("9002300043.0"))
	Info(AmountToCny("12345678901234567.781"))

}

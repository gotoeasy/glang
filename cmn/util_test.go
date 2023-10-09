package cmn

import (
	"math"
	"testing"
)

func Test_chk(t *testing.T) {
	Info(IsEmail("1@1.1"))
	Info(NormalizeVer("v1.2.3"))
}
func Test_uid(t *testing.T) {
	for i := 0; i < 10; i++ {
		Info(ULID())
	}
}

func Test_str(t *testing.T) {
	Info(Float64ToStringRoundDown(1234567890123456, 2))
	Info(Float64ToStringRoundDown(999999999999.999, 2))
	Info(Float64ToCny(1234567890123456.789))
	Info(Float64ToCny(-1234567890123456.789))
	Info(Float64ToCny(-9234567890123456.789))
	Info(Float64ToCny(math.MaxInt32 + 0.789))
	Info(Float64ToCny(math.MaxUint32 + 0.789))
	Info(Float64ToCny(999999999999.999))
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

func Test_CamelToUnderline(t *testing.T) {

	Info(CamelToUnderline("dsAdddSSdsdsAA"))
	Info(UnderlineToCamel("ds_Addd_SSdsds_AA"))

}

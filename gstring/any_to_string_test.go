package gstring

import (
	"log"
	"math"
	"testing"
)

func Test_tostring(t *testing.T) {
	var v uint64 = math.MaxUint64
	log.Println(ToString(v))

	// bt := ToBytes("sdsa刷刷")
	ss := PadLeft("bt", "sssssasss", 3)
	log.Println(ss)
}

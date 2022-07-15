package gstring

import (
	"glang/gtype"
	"log"
	"math"
	"testing"
)

func Test_tostring(t *testing.T) {
	var v uint64 = math.MaxUint64
	log.Println(ToString(v))

	bt := ToBytes("sdsa刷刷")
	log.Println(gtype.AnyType(bt))
}

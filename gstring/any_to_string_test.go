package gstring

import (
	"log"
	"math"
	"testing"

	"github.com/gotoeasy/glang/gtype"
)

func Test_tostring(t *testing.T) {
	var v uint64 = math.MaxUint64
	log.Println(ToString(v))

	bt := ToBytes("sdsa刷刷")
	log.Println(gtype.AnyType(bt))
}

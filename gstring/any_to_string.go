package gstring

import (
	"unsafe"

	"github.com/gotoeasy/glang/gconv"
)

func ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func ToString(v any, defaults ...string) string {
	s, err := gconv.AnyToString(v)
	if err != nil && len(defaults) > 0 {
		return defaults[0]
	}
	return s
}

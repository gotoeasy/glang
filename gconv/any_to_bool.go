package gconv

import (
	"errors"
	"strconv"

	"github.com/gotoeasy/glang/gtype"
)

func ToBool(v any, defaults ...bool) bool {
	rs, err := anyToBool(v)
	if err != nil && len(defaults) > 0 {
		return defaults[0]
	}
	return rs
}

func anyToBool(v any) (bool, error) {
	v = gtype.AnyType(v)

	switch b := v.(type) {
	case bool:
		return b, nil
	case string:
		return strconv.ParseBool(v.(string))
	default:
		return false, errors.New("转换失败")
	}
}

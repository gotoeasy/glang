package gconv

import (
	"errors"
	"strconv"

	"github.com/gotoeasy/glang/gtype"
)

func ToInt(v any, defaults ...int) int {
	rs, err := anyToInt(v)
	if err != nil && len(defaults) > 0 {
		return defaults[0]
	}
	return rs
}

func anyToInt(v any) (int, error) {
	v = gtype.AnyType(v)

	switch s := v.(type) {
	case int64:
		return int(s), nil
	case int32:
		return int(s), nil
	case int16:
		return int(s), nil
	case int8:
		return int(s), nil
	case uint:
		return int(s), nil
	case uint64:
		return int(s), nil
	case uint32:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint8:
		return int(s), nil
	case float64:
		return int(s), nil
	case float32:
		return int(s), nil
	case string:
		v, err := strconv.ParseInt(s, 0, 0)
		if err == nil {
			return int(v), nil
		}
		return 0, errors.New("转换失败")
	case bool:
		if s {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, nil
	default:
		return 0, errors.New("转换失败")
	}
}

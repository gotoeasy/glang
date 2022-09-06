package gconv

import (
	"errors"
	"strconv"

	"github.com/gotoeasy/glang/gtype"
)

func ToString(v any, defaults ...string) string {
	rs, err := AnyToString(v)
	if err != nil && len(defaults) > 0 {
		return defaults[0]
	}
	return rs
}

func AnyToString(v any) (string, error) {
	v = gtype.AnyType(v)

	switch s := v.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case nil:
		return "", nil
	case error:
		return s.Error(), nil
	case []byte:
		return string(v.([]byte)), nil
	default:
		return "", errors.New("转换失败")
	}
}

package cmn

import (
	"reflect"
	"time"
)

type DbValue struct {
	value any
}

func (d *DbValue) String() string {
	if d.value == nil {
		return ""
	}

	switch reflect.ValueOf(d.value).Kind() {
	case reflect.String:
		return d.value.(string)
	case reflect.Int:
		i := d.value.(int)
		return Int64ToString(int64(i))
	case reflect.Int8:
		i := d.value.(int8)
		return Int64ToString(int64(i))
	case reflect.Int16:
		i := d.value.(int16)
		return Int64ToString(int64(i))
	case reflect.Int32:
		i := d.value.(int32)
		return Int64ToString(int64(i))
	case reflect.Int64:
		i := d.value.(int64)
		return Int64ToString(i)
	case reflect.Uint:
		i := d.value.(uint)
		return Int64ToString(int64(i))
	case reflect.Uint8:
		i := d.value.(uint8)
		return Int64ToString(int64(i))
	case reflect.Uint16:
		i := d.value.(uint16)
		return Int64ToString(int64(i))
	case reflect.Uint32:
		i := d.value.(uint32)
		return Uint32ToString(i)
	case reflect.Uint64:
		i := d.value.(uint64)
		return Uint64ToString(i)
	case reflect.Float32:
		i := d.value.(float32)
		return Float64ToString(float64(i))
	case reflect.Float64:
		i := d.value.(float64)
		return Float64ToString(i)
	default:
		if b, ok := d.value.([]uint8); ok {
			return string(b)
		}
	}
	return ""
}

func (d *DbValue) Int() int {
	return StringToInt(Int64ToString(d.Int64()), 0)
}

func (d *DbValue) Int64() int64 {
	if d.value == nil {
		return 0
	}

	switch reflect.ValueOf(d.value).Kind() {
	case reflect.Int:
		i := d.value.(int)
		return int64(i)
	case reflect.Int8:
		i := d.value.(int8)
		return int64(i)
	case reflect.Int16:
		i := d.value.(int16)
		return int64(i)
	case reflect.Int32:
		i := d.value.(int32)
		return int64(i)
	case reflect.Int64:
		i := d.value.(int64)
		return i
	case reflect.Uint:
		i := d.value.(uint)
		return int64(i)
	case reflect.Uint8:
		i := d.value.(uint8)
		return int64(i)
	case reflect.Uint16:
		i := d.value.(uint16)
		return int64(i)
	case reflect.Uint32:
		i := d.value.(uint32)
		return int64(i)
	case reflect.Uint64:
		i := d.value.(uint64)
		return int64(i)
	case reflect.Float32:
		i := d.value.(float32)
		return Float64ToInt64(float64(i))
	case reflect.Float64:
		i := d.value.(float64)
		return Float64ToInt64(i)
	case reflect.String:
		s := d.value.(string)
		return StringToInt64(s, 0)
	default:
		if b, ok := d.value.([]uint8); ok {
			s := string(b)
			return StringToInt64(s, 0)
		}
	}
	return 0
}

func (d *DbValue) Float64() float64 {
	if d.value == nil {
		return 0
	}

	switch reflect.ValueOf(d.value).Kind() {
	case reflect.Int:
		i := d.value.(int)
		return float64(i)
	case reflect.Int8:
		i := d.value.(int8)
		return float64(i)
	case reflect.Int16:
		i := d.value.(int16)
		return float64(i)
	case reflect.Int32:
		i := d.value.(int32)
		return float64(i)
	case reflect.Int64:
		i := d.value.(int64)
		return float64(i)
	case reflect.Uint:
		i := d.value.(uint)
		return float64(i)
	case reflect.Uint8:
		i := d.value.(uint8)
		return float64(i)
	case reflect.Uint16:
		i := d.value.(uint16)
		return float64(i)
	case reflect.Uint32:
		i := d.value.(uint32)
		return float64(i)
	case reflect.Uint64:
		i := d.value.(uint64)
		return float64(i)
	case reflect.Float32:
		i := d.value.(float32)
		return float64(i)
	case reflect.Float64:
		i := d.value.(float64)
		return i
	case reflect.String:
		s := d.value.(string)
		return String2Float64(s, 0)
	default:
		if b, ok := d.value.([]uint8); ok {
			s := string(b)
			return String2Float64(s, 0)
		}
	}
	return 0
}

func (d *DbValue) Time() time.Time {
	if d.value == nil {
		var t time.Time
		return t
	}

	timeLayout := "2006-01-02 15:04:05"

	typ := reflect.TypeOf(d.value).String()
	if typ == "string" || typ == "[]uint8" {
		t, _ := time.ParseInLocation(timeLayout, d.String(), time.Local)
		return t
	}
	if typ == "time.Time" {
		return d.value.(time.Time)
	}
	// if typ == "ora.TimeStamp" {
	// 	return time.Time(v.(ora.TimeStamp))
	// }
	return time.Time{}
}

func (d *DbValue) RawValue() any {
	return d.value
}

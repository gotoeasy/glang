package cmn

import (
	"errors"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

// 转数据库表名（驼峰转下划线小写）
func DbTableName(name string) string {
	return CamelToUnderline(name)
}

// 转数据库字段名（驼峰转下划线小写）
func DbFieldName(name string) string {
	return CamelToUnderline(name)
}

// 反射解析出结构体类型
// 参数仅接受：结构体对象切片指针
func GetTypeOfStructSlicePointer(structSlicePointer any) (reflect.Type, error) {
	value := reflect.ValueOf(structSlicePointer)
	if value.Kind() != reflect.Ptr {
		return nil, errors.New("参数类型有误，仅支持结构体或结构体切片的指针")
	}

	typ := reflect.TypeOf(value.Elem().Interface())
	if typ.Kind() != reflect.Slice {
		return nil, errors.New("参数类型有误，仅支持结构体或结构体切片的指针")
	}

	typElem := typ.Elem()
	if typElem.Kind() != reflect.Struct {
		return nil, errors.New("参数类型有误，仅支持结构体或结构体切片的指针")
	}

	return typElem, nil
}

// 反射解析出结构体类型
// 参数仅接受：结构体对象指针
func GetTypeOfStructPointer(structPointer any) (reflect.Type, error) {
	value := reflect.ValueOf(structPointer)
	if value.Kind() != reflect.Ptr {
		return nil, errors.New("参数类型有误，仅支持结构体的指针")
	}

	typ := reflect.TypeOf(value.Elem().Interface())
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("参数类型有误，仅支持结构体的指针")
	}

	return typ, nil
}

// 反射解析出结构体类型
// 参数仅接受：结构体对象、结构体对象指针、结构体对象切片、结构体对象切片指针
func ParseStructType(obj any) (reflect.Type, error) {
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Ptr {
		return ParseStructType(value.Elem().Interface())
	}

	typ := reflect.TypeOf(obj)
	if value.Kind() == reflect.Slice {
		typ = typ.Elem()
		if typ.Kind() != reflect.Struct {
			return nil, errors.New("参数类型有误，仅支持结构体或结构体切片及其指针")
		}
	} else if value.Kind() != reflect.Struct {
		return nil, errors.New("参数类型有误，仅支持结构体或结构体切片及其指针")
	}

	return typ, nil
}

func ParseStructFieldType(structType reflect.Type) map[string]string {
	mapRs := make(map[string]string)
	if structType.Kind() == reflect.Struct {
		for i := 0; i < structType.NumField(); i++ {
			n := structType.Field(i).Name
			s := structType.Field(i).Type.String()
			mapRs[n] = s
			mapRs[DbFieldName(n)] = s
		}
	}
	return mapRs
}

func StructToMap(obj any) map[string]any {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	field := t.NumField()
	var m = make(map[string]any)
	for i := 0; field > i; i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}

func MapToStruct(mapFrom map[string]any, objPtrTo any) {
	mapstructure.Decode(mapFrom, objPtrTo)
}

func DbValue2String(v any) string {
	switch reflect.ValueOf(v).Kind() {
	case reflect.String:
		return v.(string)
	case reflect.Int:
		i := v.(int)
		return Int64ToString(int64(i))
	case reflect.Int8:
		i := v.(int8)
		return Int64ToString(int64(i))
	case reflect.Int16:
		i := v.(int16)
		return Int64ToString(int64(i))
	case reflect.Int32:
		i := v.(int32)
		return Int64ToString(int64(i))
	case reflect.Int64:
		i := v.(int64)
		return Int64ToString(i)
	case reflect.Uint:
		i := v.(uint)
		return Int64ToString(int64(i))
	case reflect.Uint8:
		i := v.(uint8)
		return Int64ToString(int64(i))
	case reflect.Uint16:
		i := v.(uint16)
		return Int64ToString(int64(i))
	case reflect.Uint32:
		i := v.(uint32)
		return Uint32ToString(i)
	case reflect.Uint64:
		i := v.(uint64)
		return Uint64ToString(i)
	case reflect.Float32:
		i := v.(float32)
		return Float64ToString(float64(i))
	case reflect.Float64:
		i := v.(float64)
		return Float64ToString(i)
	default:
		if b, ok := v.([]uint8); ok {
			return string(b)
		}
	}
	return ""
}

func DbValue2Int(v any) int {
	return StringToInt(Int64ToString(DbValue2Int64(v)), 0)
}

func DbValue2Int64(v any) int64 {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int:
		i := v.(int)
		return int64(i)
	case reflect.Int8:
		i := v.(int8)
		return int64(i)
	case reflect.Int16:
		i := v.(int16)
		return int64(i)
	case reflect.Int32:
		i := v.(int32)
		return int64(i)
	case reflect.Int64:
		i := v.(int64)
		return i
	case reflect.Uint:
		i := v.(uint)
		return int64(i)
	case reflect.Uint8:
		i := v.(uint8)
		return int64(i)
	case reflect.Uint16:
		i := v.(uint16)
		return int64(i)
	case reflect.Uint32:
		i := v.(uint32)
		return int64(i)
	case reflect.Uint64:
		i := v.(uint64)
		return int64(i)
	case reflect.Float32:
		i := v.(float32)
		return Float64ToInt64(float64(i))
	case reflect.Float64:
		i := v.(float64)
		return Float64ToInt64(i)
	case reflect.String:
		s := v.(string)
		return StringToInt64(s, 0)
	default:
		if b, ok := v.([]uint8); ok {
			s := string(b)
			return StringToInt64(s, 0)
		}
	}
	return 0
}

func DbValue2Float64(v any) float64 {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int:
		i := v.(int)
		return float64(i)
	case reflect.Int8:
		i := v.(int8)
		return float64(i)
	case reflect.Int16:
		i := v.(int16)
		return float64(i)
	case reflect.Int32:
		i := v.(int32)
		return float64(i)
	case reflect.Int64:
		i := v.(int64)
		return float64(i)
	case reflect.Uint:
		i := v.(uint)
		return float64(i)
	case reflect.Uint8:
		i := v.(uint8)
		return float64(i)
	case reflect.Uint16:
		i := v.(uint16)
		return float64(i)
	case reflect.Uint32:
		i := v.(uint32)
		return float64(i)
	case reflect.Uint64:
		i := v.(uint64)
		return float64(i)
	case reflect.Float32:
		i := v.(float32)
		return float64(i)
	case reflect.Float64:
		i := v.(float64)
		return i
	case reflect.String:
		s := v.(string)
		return String2Float64(s, 0)
	default:
		if b, ok := v.([]uint8); ok {
			s := string(b)
			return String2Float64(s, 0)
		}
	}
	return 0
}

func DbValue2Time(v any) time.Time {
	timeLayout := "2006-01-02 15:04:05"

	typ := reflect.TypeOf(v).String()
	if typ == "string" || typ == "[]uint8" {
		t, _ := time.ParseInLocation(timeLayout, DbValue2String(v), time.Local)
		return t
	}
	if typ == "time.Time" {
		return v.(time.Time)
	}
	// if typ == "ora.TimeStamp" {
	// 	return time.Time(v.(ora.TimeStamp))
	// }
	return time.Time{}
}

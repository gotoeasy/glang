package db

import (
	"errors"
	"reflect"

	"github.com/gotoeasy/glang/cmn"
	"github.com/mitchellh/mapstructure"
)

// 转数据库表名（驼峰转下划线小写）
func DbTableName(name string) string {
	return cmn.CamelToUnderline(name)
}

// 转数据库字段名（驼峰转下划线小写）
func DbFieldName(name string) string {
	return cmn.CamelToUnderline(name)
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

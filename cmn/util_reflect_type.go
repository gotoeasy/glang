package cmn

import "reflect"

// 判断类型是否为bool
func IsTypeOfBool(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Bool
}

// 判断类型是否为int
func IsTypeOfInt(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Int
}

// 判断类型是否为int8
func IsTypeOfInt8(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Int8
}

// 判断类型是否为int16
func IsTypeOfInt16(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Int16
}

// 判断类型是否为int32
func IsTypeOfInt32(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Int32
}

// 判断类型是否为int64
func IsTypeOfInt64(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Int64
}

// 判断类型是否为uint
func IsTypeOfUint(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Uint
}

// 判断类型是否为uint8
func IsTypeOfUint8(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Uint8
}

// 判断类型是否为uint16
func IsTypeOfUint16(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Uint16
}

// 判断类型是否为uint32
func IsTypeOfUint32(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Uint32
}

// 判断类型是否为uint64
func IsTypeOfUint64(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Uint64
}

// 判断类型是否为uintptr
func IsTypeOfUintptr(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Uintptr
}

// 判断类型是否为float32
func IsTypeOfFloat32(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Float32
}

// 判断类型是否为float64
func IsTypeOfFloat64(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Float64
}

// 判断类型是否为complex64
func IsTypeOfComplex64(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Complex64
}

// 判断类型是否为complex128
func IsTypeOfComplex128(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Complex128
}

// 判断类型是否为array
func IsTypeOfArray(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Array
}

// 判断类型是否为chan
func IsTypeOfChan(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Chan
}

// 判断类型是否为func
func IsTypeOfFunc(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Func
}

// 判断类型是否为interface
func IsTypeOfInterface(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Interface
}

// 判断类型是否为map
func IsTypeOfMap(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Map
}

// 判断类型是否为pointer
func IsTypeOfPointer(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Pointer
}

// 判断类型是否为slice
func IsTypeOfSlice(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Slice
}

// 判断类型是否为string
func IsTypeOfString(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.String
}

// 判断类型是否为struct
func IsTypeOfStruct(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Struct
}

// 判断类型是否为unsafePointer
func IsTypeOfUnsafePointer(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.UnsafePointer
}

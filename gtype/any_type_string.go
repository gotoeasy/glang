package gtype

import (
	"fmt"
	"reflect"
)

func AnyType(v any) any {
	if v == nil {
		return nil
	}

	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	var fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	o := reflect.ValueOf(v)
	for !o.Type().Implements(fmtStringerType) && !o.Type().Implements(errorType) && o.Kind() == reflect.Ptr && !o.IsNil() {
		o = o.Elem()
	}
	return o.Interface()
}

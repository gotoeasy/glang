package db

import (
	"bytes"
	"reflect"

	"github.com/gotoeasy/glang/cmn"
)

type SqlInserter struct {
	params []any
	entity any
	buffer *bytes.Buffer
}

func NewSqlInserter() *SqlInserter {
	return &SqlInserter{
		buffer: new(bytes.Buffer),
	}
}

func (o *SqlInserter) Insert(entity any) *SqlInserter {
	value := reflect.ValueOf(entity)
	if value.Kind() == reflect.Ptr {
		return o.Insert(value.Elem().Interface())
	} else if value.Kind() == reflect.Struct {
		// TODO 设定 CreateUser,CreateTime
		o.entity = entity
	} else {
		cmn.Error("实体类型有误，Insert仅支持结构体对象参数", entity)
	}
	return o
}

func (o *SqlInserter) GetSql() string {
	if o.entity == nil {
		return ""
	}

	typ := reflect.TypeOf(o.entity)
	value := reflect.ValueOf(o.entity)
	var params []any
	columns := ""
	values := ""
	for i := 0; i < typ.NumField(); i++ {
		if columns != "" {
			columns += ","
			values += ","
		}
		columns += DbFieldName(typ.Field(i).Name)
		values += "?"
		params = append(params, value.Field(i).Interface())
	}

	sql := "INSERT INTO " + DbTableName(value.Type().Name()) + " (" + columns + ") VALUES (" + values + ")"
	o.params = params

	return sql
}

func (o *SqlInserter) GetParameters() []any {
	return o.params
}

func (o *SqlInserter) Build() (string, []any) {
	return o.GetSql(), o.params
}

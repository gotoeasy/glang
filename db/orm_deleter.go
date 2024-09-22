package cmn

import (
	"bytes"

	"github.com/gotoeasy/glang/cmn"
)

type SqlDeleter struct {
	table  string
	params []any
	buffer *bytes.Buffer
}

func NewSqlDeleter() *SqlDeleter {
	return &SqlDeleter{
		buffer: new(bytes.Buffer),
	}
}

func (d *SqlDeleter) Delete(tableName string) *SqlDeleter {
	d.table = tableName
	return d
}

func (d *SqlDeleter) Where(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " = ", param)
}

func (d *SqlDeleter) And(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " = ", param)
}

func (d *SqlDeleter) Or(column string, param any) *SqlDeleter {
	return d.condition(" OR ", column, " = ", param)
}

func (d *SqlDeleter) Eq(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " = ", param)
}

func (d *SqlDeleter) OrEq(column string, param any) *SqlDeleter {
	return d.condition(" OR ", column, " = ", param)
}

func (d *SqlDeleter) NotEq(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " <> ", param)
}

func (d *SqlDeleter) OrNotEq(column string, param any) *SqlDeleter {
	return d.condition(" OR ", column, " <> ", param)
}

func (d *SqlDeleter) Gt(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " > ", param)
}

func (d *SqlDeleter) OrGt(column string, param any) *SqlDeleter {
	return d.condition(" OR ", column, " > ", param)
}

func (d *SqlDeleter) Ge(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " >= ", param)
}

func (d *SqlDeleter) OrGe(column string, param any) *SqlDeleter {
	return d.condition(" OR ", column, " >= ", param)
}

func (d *SqlDeleter) Lt(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " < ", param)
}

func (d *SqlDeleter) OrLt(column string, param any) *SqlDeleter {
	return d.condition(" OR ", column, " < ", param)
}

func (d *SqlDeleter) Le(column string, param any) *SqlDeleter {
	return d.condition(" AND ", column, " <= ", param)
}

func (d *SqlDeleter) OrLe(column string, param any) *SqlDeleter {
	return d.condition(" OR ", column, " <= ", param)
}

func (d *SqlDeleter) Like(column string, param string) *SqlDeleter {
	return d.condition(" AND ", column, " LIKE ", "%"+param+"%")
}

func (d *SqlDeleter) NotLike(column string, param string) *SqlDeleter {
	return d.condition(" AND ", column, " NOT LIKE ", "%"+param+"%")
}

func (d *SqlDeleter) LeftLike(column string, param string) *SqlDeleter {
	return d.condition(" AND ", column, " LIKE ", param+"%")
}

func (d *SqlDeleter) RightLike(column string, param string) *SqlDeleter {
	return d.condition(" AND ", column, " LIKE ", "%"+param)
}

func (d *SqlDeleter) In(column string, params ...any) *SqlDeleter {
	if len(params) < 1 {
		return d
	}

	in := ""
	for _, v := range params {
		if cmn.Len(in) > 0 {
			in += ", "
		}
		in += "?"
		d.params = append(d.params, v)
	}

	if d.buffer.Len() > 0 {
		d.buffer.WriteString(" AND ")
	}
	d.buffer.WriteString(column + " IN (" + in + ")")
	return d
}

func (d *SqlDeleter) NotIn(column string, params ...any) *SqlDeleter {
	if len(params) < 1 {
		return d
	}

	in := ""
	for _, v := range params {
		if cmn.Len(in) > 0 {
			in += ", "
		}
		in += "?"
		d.params = append(d.params, v)
	}

	if d.buffer.Len() > 0 {
		d.buffer.WriteString(" AND ")
	}
	d.buffer.WriteString(column + " NOT IN (" + in + ")")
	return d
}

func (d *SqlDeleter) IsNull(column string) *SqlDeleter {
	if d.buffer.Len() > 0 {
		d.buffer.WriteString(" AND ")
	}
	d.buffer.WriteString(column + " IS NULL")
	return d
}

func (d *SqlDeleter) IsNotNull(column string) *SqlDeleter {
	if d.buffer.Len() > 0 {
		d.buffer.WriteString(" AND ")
	}
	d.buffer.WriteString(column + " IS NOT NULL")
	return d
}

func (d *SqlDeleter) condition(exprAndOr string, column string, comparer string, param any) *SqlDeleter {
	if d.buffer.Len() > 0 {
		d.buffer.WriteString(exprAndOr)
	}
	d.buffer.WriteString(column)
	d.buffer.WriteString(comparer)
	d.buffer.WriteString("?")

	d.params = append(d.params, param)
	return d
}

func (d *SqlDeleter) GetSql() string {
	return "DELETE FROM " + d.table + " WHERE " + d.buffer.String()
}

func (d *SqlDeleter) GetParameters() []any {
	return d.params
}

func (d *SqlDeleter) Build() (string, []any) {
	return d.GetSql(), d.params
}

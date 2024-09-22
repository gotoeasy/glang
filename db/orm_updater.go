package cmn

import (
	"bytes"

	"github.com/gotoeasy/glang/cmn"
)

type SqlUpdater struct {
	table     string
	params    []any
	bufferSet *bytes.Buffer
	buffer    *bytes.Buffer
}

func NewSqlUpdater() *SqlUpdater {
	return &SqlUpdater{
		bufferSet: new(bytes.Buffer),
		buffer:    new(bytes.Buffer),
	}
}

func (d *SqlUpdater) Update(tableName string) *SqlUpdater {
	d.table = tableName
	return d
}

func (d *SqlUpdater) SetMap(m map[string]any) *SqlUpdater {
	for column, param := range m {
		if d.bufferSet.Len() > 0 {
			d.bufferSet.WriteString(", ")
		}

		if param == nil {
			d.bufferSet.WriteString(column + " = NULL")
		} else {
			d.bufferSet.WriteString(column + " = ?")
			d.params = append(d.params, param)
		}
	}

	return d
}

func (d *SqlUpdater) Set(column string, param any) *SqlUpdater {
	if d.bufferSet.Len() > 0 {
		d.bufferSet.WriteString(", ")
	}
	d.bufferSet.WriteString(column + " = ?")
	d.params = append(d.params, param)
	return d
}

func (d *SqlUpdater) Where(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " = ", param)
}

func (d *SqlUpdater) And(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " = ", param)
}

func (d *SqlUpdater) Or(column string, param any) *SqlUpdater {
	return d.condition(" OR ", column, " = ", param)
}

func (d *SqlUpdater) Eq(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " = ", param)
}

func (d *SqlUpdater) OrEq(column string, param any) *SqlUpdater {
	return d.condition(" OR ", column, " = ", param)
}

func (d *SqlUpdater) NotEq(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " <> ", param)
}

func (d *SqlUpdater) OrNotEq(column string, param any) *SqlUpdater {
	return d.condition(" OR ", column, " <> ", param)
}

func (d *SqlUpdater) Gt(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " > ", param)
}

func (d *SqlUpdater) OrGt(column string, param any) *SqlUpdater {
	return d.condition(" OR ", column, " > ", param)
}

func (d *SqlUpdater) Ge(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " >= ", param)
}

func (d *SqlUpdater) OrGe(column string, param any) *SqlUpdater {
	return d.condition(" OR ", column, " >= ", param)
}

func (d *SqlUpdater) Lt(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " < ", param)
}

func (d *SqlUpdater) OrLt(column string, param any) *SqlUpdater {
	return d.condition(" OR ", column, " < ", param)
}

func (d *SqlUpdater) Le(column string, param any) *SqlUpdater {
	return d.condition(" AND ", column, " <= ", param)
}

func (d *SqlUpdater) OrLe(column string, param any) *SqlUpdater {
	return d.condition(" OR ", column, " <= ", param)
}

func (d *SqlUpdater) Like(column string, param string) *SqlUpdater {
	return d.condition(" AND ", column, " LIKE ", "%"+param+"%")
}

func (d *SqlUpdater) NotLike(column string, param string) *SqlUpdater {
	return d.condition(" AND ", column, " NOT LIKE ", "%"+param+"%")
}

func (d *SqlUpdater) LeftLike(column string, param string) *SqlUpdater {
	return d.condition(" AND ", column, " LIKE ", param+"%")
}

func (d *SqlUpdater) RightLike(column string, param string) *SqlUpdater {
	return d.condition(" AND ", column, " LIKE ", "%"+param)
}

func (d *SqlUpdater) In(column string, params ...any) *SqlUpdater {
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

func (d *SqlUpdater) NotIn(column string, params ...any) *SqlUpdater {
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

func (d *SqlUpdater) IsNull(column string) *SqlUpdater {
	if d.buffer.Len() > 0 {
		d.buffer.WriteString(" AND ")
	}
	d.buffer.WriteString(column + " IS NULL")
	return d
}

func (d *SqlUpdater) IsNotNull(column string) *SqlUpdater {
	if d.buffer.Len() > 0 {
		d.buffer.WriteString(" AND ")
	}
	d.buffer.WriteString(column + " IS NOT NULL")
	return d
}

func (d *SqlUpdater) condition(exprAndOr string, column string, comparer string, param any) *SqlUpdater {
	if d.buffer.Len() > 0 {
		d.buffer.WriteString(exprAndOr)
	}
	d.buffer.WriteString(column)
	d.buffer.WriteString(comparer)
	d.buffer.WriteString("?")

	d.params = append(d.params, param)
	return d
}

func (d *SqlUpdater) GetSql() string {
	return "UPDATE " + d.table + " SET " + d.bufferSet.String() + " WHERE " + d.buffer.String()
}

func (d *SqlUpdater) GetParameters() []any {
	return d.params
}

func (d *SqlUpdater) Build() (string, []any) {
	return d.GetSql(), d.params
}

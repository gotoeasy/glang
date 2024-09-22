package db

import (
	"bytes"
	"strings"

	"github.com/gotoeasy/glang/cmn"
)

type SqlSelector struct {
	params     []any
	bufSelect  *bytes.Buffer
	bufFrom    *bytes.Buffer
	bufWhere   *bytes.Buffer
	bufOrder   *bytes.Buffer
	bufCurrent *bytes.Buffer
}

func NewSqlSelector() *SqlSelector {
	return &SqlSelector{
		bufSelect: new(bytes.Buffer),
		bufFrom:   new(bytes.Buffer),
		bufWhere:  new(bytes.Buffer),
		bufOrder:  new(bytes.Buffer),
	}
}

func (d *SqlSelector) Select(columns ...string) *SqlSelector {
	d.bufCurrent = d.bufSelect
	if len(columns) > 0 {
		d.bufCurrent.WriteString(strings.Join(columns, ","))
	}
	return d
}

func (d *SqlSelector) From(table string, alias ...string) *SqlSelector {
	d.bufCurrent = d.bufFrom
	if d.bufCurrent.Len() > 0 {
		d.bufCurrent.WriteString(", ")
	}
	d.bufCurrent.WriteString(table)
	if len(alias) > 0 {
		d.bufCurrent.WriteString(" " + alias[0])
	}
	return d
}

func (d *SqlSelector) Join(table string, alias ...string) *SqlSelector {
	d.bufCurrent = d.bufFrom
	d.bufCurrent.WriteString(" JOIN ")
	d.bufCurrent.WriteString(table)
	if len(alias) > 0 {
		d.bufCurrent.WriteString(" " + alias[0])
	}
	return d
}

func (d *SqlSelector) On(onCond string) *SqlSelector {
	d.bufCurrent = d.bufFrom
	d.bufCurrent.WriteString(" ON ")
	d.bufCurrent.WriteString(onCond)
	return d
}

func (d *SqlSelector) And(cond string, params ...any) *SqlSelector {
	if d.bufCurrent.Len() > 0 {
		d.bufCurrent.WriteString(" AND ")
	}
	d.bufCurrent.WriteString(cond)
	if len(params) > 0 {
		d.params = append(d.params, params...)
	}
	return d
}

func (d *SqlSelector) Or(cond string, params ...any) *SqlSelector {
	if d.bufCurrent.Len() > 0 {
		d.bufCurrent.WriteString(" OR ")
	}
	d.bufCurrent.WriteString(cond)
	if len(params) > 0 {
		d.params = append(d.params, params...)
	}
	return d
}

func (d *SqlSelector) Where(cond string, params ...any) *SqlSelector {
	d.bufCurrent = d.bufWhere
	if d.bufCurrent.Len() > 0 {
		d.bufCurrent.WriteString(" AND ")
	} else {
		d.bufCurrent.WriteString(" WHERE ")
	}
	d.bufCurrent.WriteString(cond)
	if len(params) > 0 {
		d.params = append(d.params, params...)
	}
	return d
}

func (d *SqlSelector) OrderBy(orders ...string) *SqlSelector {
	if len(orders) > 0 {
		d.bufOrder.WriteString(strings.Join(orders, ","))
	}
	return d
}

func (d *SqlSelector) Eq(column string, param any) *SqlSelector {
	return d.condition(" AND ", column, " = ", param)
}

func (d *SqlSelector) OrEq(column string, param any) *SqlSelector {
	return d.condition(" OR ", column, " = ", param)
}

func (d *SqlSelector) NotEq(column string, param any) *SqlSelector {
	return d.condition(" AND ", column, " <> ", param)
}

func (d *SqlSelector) OrNotEq(column string, param any) *SqlSelector {
	return d.condition(" OR ", column, " <> ", param)
}

func (d *SqlSelector) Gt(column string, param any) *SqlSelector {
	return d.condition(" AND ", column, " > ", param)
}

func (d *SqlSelector) OrGt(column string, param any) *SqlSelector {
	return d.condition(" OR ", column, " > ", param)
}

func (d *SqlSelector) Ge(column string, param any) *SqlSelector {
	return d.condition(" AND ", column, " >= ", param)
}

func (d *SqlSelector) OrGe(column string, param any) *SqlSelector {
	return d.condition(" OR ", column, " >= ", param)
}

func (d *SqlSelector) Lt(column string, param any) *SqlSelector {
	return d.condition(" AND ", column, " < ", param)
}

func (d *SqlSelector) OrLt(column string, param any) *SqlSelector {
	return d.condition(" OR ", column, " < ", param)
}

func (d *SqlSelector) Le(column string, param any) *SqlSelector {
	return d.condition(" AND ", column, " <= ", param)
}

func (d *SqlSelector) OrLe(column string, param any) *SqlSelector {
	return d.condition(" OR ", column, " <= ", param)
}

func (d *SqlSelector) Like(column string, param string) *SqlSelector {
	return d.condition(" AND ", column, " LIKE ", "%"+param+"%")
}

func (d *SqlSelector) NotLike(column string, param string) *SqlSelector {
	return d.condition(" AND ", column, " NOT LIKE ", "%"+param+"%")
}

func (d *SqlSelector) LeftLike(column string, param string) *SqlSelector {
	return d.condition(" AND ", column, " LIKE ", param+"%")
}

func (d *SqlSelector) RightLike(column string, param string) *SqlSelector {
	return d.condition(" AND ", column, " LIKE ", "%"+param)
}

func (d *SqlSelector) In(column string, params ...any) *SqlSelector {
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

	if d.bufWhere.Len() > 0 {
		d.bufWhere.WriteString(" AND ")
	}
	d.bufWhere.WriteString(column + " IN (" + in + ")")
	return d
}

func (d *SqlSelector) NotIn(column string, params ...any) *SqlSelector {
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

	if d.bufWhere.Len() > 0 {
		d.bufWhere.WriteString(" AND ")
	}
	d.bufWhere.WriteString(column + " NOT IN (" + in + ")")
	return d
}

func (d *SqlSelector) IsNull(column string) *SqlSelector {
	if d.bufWhere.Len() > 0 {
		d.bufWhere.WriteString(" AND ")
	}
	d.bufWhere.WriteString(column + " IS NULL")
	return d
}

func (d *SqlSelector) IsNotNull(column string) *SqlSelector {
	if d.bufWhere.Len() > 0 {
		d.bufWhere.WriteString(" AND ")
	}
	d.bufWhere.WriteString(column + " IS NOT NULL")
	return d
}

func (d *SqlSelector) condition(exprAndOr string, column string, comparer string, param any) *SqlSelector {
	if d.bufCurrent.Len() > 0 {
		d.bufCurrent.WriteString(exprAndOr)
	}
	d.bufCurrent.WriteString(column)
	d.bufCurrent.WriteString(comparer)
	d.bufCurrent.WriteString("?")

	d.params = append(d.params, param)
	return d
}

func (d *SqlSelector) Append(appendSql string, params ...any) *SqlSelector {
	d.bufWhere.WriteString(appendSql)
	if len(params) > 0 {
		d.params = append(d.params, params...)
	}

	return d
}

func (d *SqlSelector) GetSql() string {
	s := "SELECT " + d.bufSelect.String() + " FROM " + d.bufFrom.String() + d.bufWhere.String()
	if d.bufOrder.Len() > 0 {
		s += " ORDER BY " + d.bufOrder.String()
	}
	return s
}

func (d *SqlSelector) GetParameters() []any {
	return d.params
}

func (d *SqlSelector) Build() (string, []any) {
	return d.GetSql(), d.params
}

package cmn

import (
	"database/sql"
	"errors"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gotoeasy/glang/cmn"
)

var _sqlDb *sql.DB

// 数据库控制器
type DbHandle struct {
	db  *sql.DB
	tx  *sql.Tx
	err error
	opt *DbOption
}

// 配置项
type DbOption struct {
	AutoCommit bool         // 是否自动提交
	GlcData    *cmn.GlcData // 需要使用GLC日志跟踪码等信息时传入
}

// 初始化数据库配置
func initSqlDb() (*sql.DB, error) {
	if _sqlDb != nil {
		return _sqlDb, nil
	}

	// TODO 读取配置
	dataSource := cmn.GetEnvStr("MySqlDataSource") // 例："user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, errors.Join(mysql.ErrInvalidConn, err)
	}

	_sqlDb = db

	_sqlDb.SetMaxIdleConns(5)                // 最大空闲连接数
	_sqlDb.SetMaxOpenConns(50)               // 最大连接数
	_sqlDb.SetConnMaxLifetime(time.Hour * 3) // 一个连接的最长生命周期(mysql默认8小时)

	return _sqlDb, nil
}

// 新建数据库控制器
func NewDbHandle(opt ...*DbOption) *DbHandle {

	db, err := initSqlDb()
	dbhandle := &DbHandle{
		db:  db, // 可能是nil
		err: err,
		opt: &DbOption{},
	}

	if len(opt) > 0 {
		dbhandle.opt = opt[0]
	}

	return dbhandle
}

// 是否自动提交
func (d *DbHandle) IsAutoCommit() bool {
	return d.opt.AutoCommit
}

// 开启事务，逐一要配对调用 defer EndTransaction()
func (d *DbHandle) BeginTransaction() {
	if d == nil {
		return // 忽视
	}
	if d.IsAutoCommit() {
		cmn.Warn("忽略开启事务操作（当前是自动提交模式）", d.opt.GlcData)
		return
	}

	if d.db == nil || d.err != nil {
		cmn.Error("开启事务失败（数据库连接无效或已发生错误）", d.err, d.opt.GlcData)
		return
	}

	if d.tx != nil {
		cmn.Error("重复开启事务（忽略操作）", d.opt.GlcData)
		return // 重复Begin不出错，简单忽视
	}

	tx, err := d.db.Begin()
	if err != nil {
		d.err = err
		cmn.Error("开启事务失败：", err, d.opt.GlcData)
	} else {
		cmn.Debug("开启事务", d.opt.GlcData)
		d.tx = tx
	}
}

// 提交事务，失败时返回错误
func (d *DbHandle) Commit() error {
	if d == nil {
		return nil // 忽视
	}
	if d.IsAutoCommit() {
		cmn.Warn("忽略提交操作（当前是自动提交模式）", d.opt.GlcData)
		return nil
	}

	if d.db == nil || d.err != nil {
		cmn.Error("提交事务失败（数据库连接无效或已发生错误）", d.err, d.opt.GlcData)
		return nil
	}

	if d.tx == nil {
		cmn.Error("无效的提交（事务还没有开始，忽略操作）", d.opt.GlcData)
		return nil
	}
	err := d.tx.Commit()
	if err == nil {
		d.tx = nil
		cmn.Debug("提交事务", d.opt.GlcData)
	} else {
		cmn.Error("提交事务失败", err, d.opt.GlcData)
	}
	return err
}

// 回滚事务，失败时返回错误
func (d *DbHandle) Rollback() error {
	if d == nil {
		return nil // 忽视
	}
	if d.IsAutoCommit() {
		cmn.Warn("忽略回滚操作（当前是自动提交模式）", d.opt.GlcData)
		return nil
	}

	if d.db == nil || d.err != nil {
		cmn.Error("回滚事务失败（数据库连接无效或已发生错误）", d.err, d.opt.GlcData)
		return nil
	}

	if d.tx == nil {
		cmn.Error("无效的回滚（事务还没有开始，忽略操作）", d.opt.GlcData)
		return nil
	}
	err := d.tx.Rollback()
	if err == nil {
		d.tx = nil
		cmn.Debug("回滚事务", d.opt.GlcData)
	} else {
		cmn.Error("回滚事务失败", err, d.opt.GlcData)
	}
	return err
}

// 结束事务（通过recover()捕获错误，没有错误时提交，否则回滚），事务操作失败时返回错误
func (d *DbHandle) EndTransaction() error {
	if e := recover(); e != nil {
		cmn.Error("panic:", e)
		if d == nil || d.IsAutoCommit() {
			return nil
		}
		return d.Rollback() // 异常时回滚
	}

	if d.IsAutoCommit() {
		return nil // 自动提交时忽略
	}

	if d == nil || d.db == nil || d.err != nil {
		return nil // NewDbHandle 或 BeginTransaction 出现问题，忽略
	}

	return d.Commit()
}

// 执行SQL（开启事务时自动在事务内执行），出错时panic
func (d *DbHandle) Execute(sql string, params ...any) int64 {
	if d.err != nil {
		panic(d.err)
	}

	if !d.IsAutoCommit() && d.tx == nil {
		d.BeginTransaction()
	}

	var cnt int64
	var err error
	if d.tx != nil {
		// 事务优先
		rs, e := d.tx.Exec(sql, params...)
		if e != nil {
			cnt, err = 0, e
		} else {
			cnt, err = rs.RowsAffected()
		}
	} else {
		// 没开启事务时直接用db
		cmn.Debug("执行SQL未使用事务", d.opt.GlcData)
		rs, e := d.db.Exec(sql, params...)
		if e != nil {
			cnt, err = 0, e
		} else {
			cnt, err = rs.RowsAffected()
		}
	}

	if err == nil {
		cmn.Info(sql, "\n  parameters:", params, "\n ", cnt, "rows affected", d.opt.GlcData)
	} else {
		cmn.Error("执行SQL发生错误：", err, "\n", sql, "\n  parameters: ", params, d.opt.GlcData)
		panic(err)
	}
	return cnt
}

// 插入记录，出错时panic
func (d *DbHandle) Insert(entity any) int64 {
	SQL, params := NewSqlInserter().Insert(entity).Build()
	return d.Execute(SQL, params...)
}

// 删除记录，出错时panic
func (d *DbHandle) Delete(deleter *SqlDeleter) int64 {
	SQL, params := deleter.Build()
	return d.Execute(SQL, params...)
}

// 更新记录，出错时panic
func (d *DbHandle) Update(updater *SqlUpdater) int64 {
	SQL, params := updater.Build()
	return d.Execute(SQL, params...)
}

// 查找返回Map切片，出错时panic
// 参数entity通常为结构体对象或其指针（仅解析用不作修改）
func (d *DbHandle) FindMaps(entity any, sql string, params ...any) []map[string]any {
	structType, err := ParseStructType(entity)
	if err != nil {
		panic(err)
	}

	mapType := ParseStructFieldType(structType)
	var maps []map[string]any

	if d.tx != nil {
		rows, err := d.tx.Query(sql, params...) // 事务优先
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, false)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	} else {
		rows, err := d.db.Query(sql, params...) // 没开启事务时直接用db
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, false)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	}

	cmn.Info(sql, "\n  parameters:", params, "\n ", len(maps), "rows selected")
	return maps
}

// 查找一条记录，返回Map或无数据时nil，出错时panic
// 参数entity通常为结构体对象或其指针（仅解析用不作修改）
func (d *DbHandle) FindMap(entity any, sql string, params ...any) map[string]any {
	structType, err := ParseStructType(entity)
	if err != nil {
		panic(err)
	}

	mapType := ParseStructFieldType(structType)
	var maps []map[string]any

	if d.tx != nil {
		rows, err := d.tx.Query(sql, params...) // 事务优先
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, true)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	} else {
		rows, err := d.db.Query(sql, params...) // 没开启事务时直接用db
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, true)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	}

	if len(maps) > 0 {
		cmn.Info(sql, "\n  parameters:", params, "\n ", 1, "rows selected", d.opt.GlcData)
		return maps[0]
	}

	cmn.Info(sql, "\n  parameters:", params, "\n ", 0, "rows selected", d.opt.GlcData)
	return nil
}

// 查找记录并存入参数所指切片，出错时panic
// 参数structSlicePtr必须是结构体对象切片的指针
func (d *DbHandle) FindList(structSlicePtr any, sql string, params ...any) {
	structType, err := GetTypeOfStructSlicePointer(structSlicePtr)
	if err != nil {
		panic(err)
	}

	mapType := ParseStructFieldType(structType)
	var maps []map[string]any

	if d.tx != nil {
		rows, err := d.tx.Query(sql, params...) // 事务优先
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, false)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	} else {
		rows, err := d.db.Query(sql, params...) // 没开启事务时直接用db
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, false)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	}

	// 查询结果存入参数所指切片
	value := reflect.ValueOf(structSlicePtr)
	root := value.Elem()
	for i := 0; i < len(maps); i++ {
		newItem := reflect.New(structType)
		obj := newItem.Interface()
		MapToStruct(maps[i], &obj)
		root.Set(reflect.Append(root, newItem.Elem()))
	}
	cmn.Info(sql, "\n  parameters:", params, "\n ", len(maps), "rows selected", d.opt.GlcData)
}

// 查找记录并存入参数所指对象，出错时panic
// 参数structPtr必须是结构体对象的指针
// 查有数据时返回true，否则false
func (d *DbHandle) FindOne(structPtr any, sql string, params ...any) bool {
	structType, err := GetTypeOfStructPointer(structPtr)
	if err != nil {
		panic(err)
	}

	mapType := ParseStructFieldType(structType)
	var maps []map[string]any

	if d.tx != nil {
		rows, err := d.tx.Query(sql, params...) // 事务优先
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, true)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	} else {
		rows, err := d.db.Query(sql, params...) // 没开启事务时直接用db
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
		maps, err = d.fetchToMapArray(mapType, rows, true)
		if err != nil {
			cmn.Error(sql, "查询出错：", err, "\n  parameters:", params, d.opt.GlcData)
			panic(err)
		}
	}

	if len(maps) > 0 {
		MapToStruct(maps[0], structPtr)
		cmn.Info(sql, "\n  parameters:", params, "\n ", 1, "rows selected", d.opt.GlcData)
		return true
	}
	cmn.Info(sql, "\n  parameters:", params, "\n ", 0, "rows selected", d.opt.GlcData)
	return false
}

// 查找件数，出错时panic
func (d *DbHandle) Count(sql string, params ...any) int {
	return d.FindInt(sql, params...)
}

// 查找单记录的单个字段值，出错时panic
func (d *DbHandle) FindInt(sql string, params ...any) int {
	return d.FindValue(sql, params...).Int()
}

// 查找单记录的单个字段值，出错时panic
func (d *DbHandle) FindFloat64(sql string, params ...any) float64 {
	return d.FindValue(sql, params...).Float64()
}

// 查找单记录的单个字段值，出错时panic
func (d *DbHandle) FindTime(sql string, params ...any) time.Time {
	return d.FindValue(sql, params...).Time()
}

// 查找单记录的单个字段值，出错时panic
func (d *DbHandle) FindString(sql string, params ...any) string {
	return d.FindValue(sql, params...).String()
}

// 查找单记录的单个字段值，出错时panic
func (d *DbHandle) FindValue(sql string, params ...any) *DbValue {
	if d.tx != nil {
		rows, err := d.tx.Query(sql, params...) // 事务优先
		if err != nil {
			panic(err)
		}

		v := d.getDbValue(rows)
		if v == nil {
			cmn.Info(sql, "\n  parameters:", params, "\n ", 0, "rows selected", d.opt.GlcData)
		} else {
			cmn.Info(sql, "\n  parameters:", params, "\n ", 1, "rows selected", d.opt.GlcData)
		}
		return v
	} else {
		rows, err := d.db.Query(sql, params...) // 没开启事务时直接用db
		if err != nil {
			panic(err)
		}
		v := d.getDbValue(rows)
		if v == nil {
			cmn.Info(sql, "\n  parameters:", params, "\n ", 0, "rows selected", d.opt.GlcData)
		} else {
			cmn.Info(sql, "\n  parameters:", params, "\n ", 1, "rows selected", d.opt.GlcData)
		}
		return v
	}
}

func (d *DbHandle) getDbValue(rows *sql.Rows) *DbValue {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	if len(columns) > 1 {
		panic(errors.New("只能查询1列，实际查询了" + cmn.IntToString(len(columns)) + "列"))
	}

	if rows.Next() {
		var v any
		err := rows.Scan(&v)
		if err != nil {
			panic(err)
		}
		return &DbValue{value: v}
	}
	return nil
}

func (d *DbHandle) fetchToMapArray(mapType map[string]string, rows *sql.Rows, once bool) ([]map[string]any, error) {
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}

	var maps []map[string]any
	for rows.Next() {
		values := make([]any, len(columns))   // 查询列等长的切片，用于存记录的各字段值
		scanArgs := make([]any, len(columns)) // 查询列等长的切片，Scan操作用的参数
		for i := range values {
			scanArgs[i] = &values[i] // 存引用
		}

		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		rowMap := make(map[string]any)
		for i, colValue := range values {
			column := columns[i]
			if colValue == nil {
				rowMap[column] = nil
				continue
			}

			dbValue := &DbValue{value: colValue}

			if mapType[column] == "string" || mapType[DbFieldName(column)] == "string" {
				rowMap[column] = dbValue.String()
			} else if mapType[column] == "time.Time" || mapType[DbFieldName(column)] == "time.Time" {
				rowMap[column] = dbValue.Time()
			} else if mapType[column] == "int" || mapType[DbFieldName(column)] == "int" {
				rowMap[column] = dbValue.Int()
			} else if mapType[column] == "float64" || mapType[DbFieldName(column)] == "float64" {
				rowMap[column] = dbValue.Float64()
			} else {
				cmn.Warn("未支持的结构体字段类型")
				rowMap[column] = dbValue
			}
		}

		maps = append(maps, rowMap)

		if once {
			break
		}
	}

	return maps, nil
}

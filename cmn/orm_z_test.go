package cmn

import (
	"testing"
	"time"
)

type xx_test struct {
	Id               string
	RoomId           string
	ReserveTitle     string
	ReserveUser      string
	ReserveStartTime time.Time
	ReserveEndTime   time.Time
	Note             string
	Status           string
	CreateUser       string
	CreateTime       time.Time
	UpdateUser       string
	UpdateTime       time.Time
}

func Test_orm_select(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			Error("出错啦", e)
		}
	}()

	dbHandle := NewDbHandle()
	var ent xx_test

	m := dbHandle.FindMap(ent, "select Id,Create_User CreateUser,Note,Update_Time UpdateTime from xx_test where id=? or 1=1", "01H1HC1ZDQH2NX87WF1715YSCX")
	Info("map----------", m)

	dbHandle.FindOne(&ent, "select Id,Create_User CreateUser,Note,Update_Time UpdateTime from xx_test where id=? or 1=1", "01H1HC1ZDQH2NX87WF1715YSCX")
	Info(ent)

	var ary []xx_test
	sqls := NewSqlSelector()
	sqls.Select("id", "Create_User CreateUser", "note", "Update_Time")
	sqls.From("xx_test")
	sqls.Where("id>?", "01H1HC1ZDQH2NX87WF1715YSCX")
	sqls.OrderBy("id desc")
	sql, params := sqls.Build()
	dbHandle.FindList(&ary, sql, params...)
	Info(ary)

	cnt := dbHandle.Count("select count(*) from xx_test")
	Info("件数", cnt)
}

func Test_orm(t *testing.T) {
	defer Recover()

	id := ULID()
	ent := &xx_test{
		Id:           id,
		RoomId:       ULID(),
		ReserveTitle: "xxxxxxxxxxxxxxxx99999",
		Note:         "测试运行完成时间: 2023/5/28 12:12:46 2",
		CreateTime:   time.Now(),
	}

	dbHandle := NewDbHandle()
	dbHandle.BeginTransaction()
	defer dbHandle.EndTransaction()

	dbHandle.Insert(ent)

	deleter := NewSqlDeleter().Delete("xx_test").Where("id", "01H1H80HGRB81P5VRAAS8KD1GH")
	dbHandle.Delete(deleter)

}

func Test_ormDel(t *testing.T) {
	dbHandle := NewDbHandle()
	dbHandle.BeginTransaction()
	defer dbHandle.EndTransaction()
	sqld := NewSqlDeleter().Delete("xx_test").Where("room_id", "1").Ge("reserve_title", 2).In("reserve_user", "a", "b", "c").Like("note", "sss")
	dbHandle.Delete(sqld)
}

func Test_ormUpd(t *testing.T) {
	dbHandle := NewDbHandle()
	dbHandle.BeginTransaction()
	defer dbHandle.EndTransaction()
	sqlu := NewSqlUpdater().Update("xx_test").Set("room_id", "111").
		Set("reserve_title", "ssssss").
		Lt("room_id", "1").Ge("reserve_title", "2").In("reserve_user", "a", "b", "c")
	dbHandle.Update(sqlu)
}

package cmn

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func Test_redis6(t *testing.T) {
	rd := NewRedis6Client(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	rd.Set("abc", "vvvvvvvvvvvvv", time.Minute)
	v, err := rd.Get("abc")
	Info(v, err)

	i := rd.Exists("abc")
	Info(i)

	rd.Del("abc")
	v, err = rd.Get("abc")
	Info(v, err)

	i2 := rd.Exists("abc")
	Info(i2)
}

func Test_leveldb(t *testing.T) {
	ldb := NewLevelDB("f:\\222\\ldbtest", nil)
	ldb.Put(StringToBytes("key"), StringToBytes("value"))

	by, err := ldb.Get(StringToBytes("key"))
	Info(BytesToString(by), err)

	// New多次也还是同一个客户端
	ldb = NewLevelDB("f:\\222\\ldbtest", nil)
	by, err = ldb.Get(StringToBytes("key"))
	Info(BytesToString(by), err)
}

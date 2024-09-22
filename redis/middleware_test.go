package redis

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gotoeasy/glang/cmn"
)

func Test_redis6(t *testing.T) {
	rd := NewRedis6Client(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	rd.Set("abc", "vvvvvvvvvvvvv", time.Minute)
	v, err := rd.Get("abc")
	cmn.Info(v, err)

	i := rd.Exists("abc")
	cmn.Info(i)

	rd.Del("abc")
	v, err = rd.Get("abc")
	cmn.Info(v, err)

	i2 := rd.Exists("abc")
	cmn.Info(i2)
}

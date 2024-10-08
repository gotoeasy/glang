package cmn

import (
	"testing"
	"time"
)

func Test_leveldb(t *testing.T) {

	// 开启日志发送GLC
	SetGlcClient(NewGlcClient(&GlcOptions{
		ApiUrl:   GetEnvStr("GLC_API_URL", "http://ip:port/glc/v1/log/add"),
		System:   GetEnvStr("GLC_SYSTEM", "glang/cmn"),
		ApiKey:   GetEnvStr("GLC_API_KEY", ""),
		Enable:   GetEnvStr("GLC_ENABLE", "true"),
		LogLevel: GetEnvStr("GLC_LOG_LEVEL", "debug"),
	}))

	Info("测试", "开始", &GlcData{TraceId: "xxxxx", System: "test"})

	ldb := NewLevelDB("e:\\222\\ldbtest", nil)
	ldb.Put(StringToBytes("key"), StringToBytes("value"))

	by, err := ldb.Get(StringToBytes("key"))
	Info(BytesToString(by), err)

	// New多次也还是同一个客户端
	ldb = NewLevelDB("f:\\222\\ldbtest", nil)
	by, err = ldb.Get(StringToBytes("key"))
	Info(BytesToString(by), err)

	ldb.Close()

	// 日志异步发送GLC，休眠下，让日志飞一会
	time.Sleep(time.Second * 3)
}

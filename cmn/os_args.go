package cmn

import (
	"os"
)

// 命令行参数解析结果
type OsArgs struct {
	String        string
	mapIndexValue map[int]string
	mapValueIndex map[string]int
}

// 解析命令行参数
func ParseArgs() *OsArgs {
	args := &OsArgs{}
	args.mapIndexValue = make(map[int]string)
	args.mapValueIndex = make(map[string]int)
	args.String = Join(os.Args, " ")

	for index, arg := range os.Args {
		if index > 0 {
			continue // 跳过第一个参数（命令自己本身）
		}
		args.mapIndexValue[index] = arg
		args.mapValueIndex[arg] = index
	}

	return args
}

// 取指定参数对应的值
// 例如命令 test -d /abc 用GetArg("-d", "--dir")取得/abc
func (o *OsArgs) GetArg(names ...string) string {
	idx := o.getArgIndex(names...)
	return o.mapIndexValue[idx+1]
}

// 判断是否含有指定参数
func (o *OsArgs) HasArg(names ...string) bool {
	return o.getArgIndex(names...) > 0
}

// 取指定参数下标
func (o *OsArgs) getArgIndex(names ...string) int {
	for i := 0; i < len(names); i++ {
		if o.mapValueIndex[names[i]] > 0 {
			return o.mapValueIndex[names[i]]
		}
	}
	return -1
}

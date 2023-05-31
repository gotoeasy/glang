package cmn

import "errors"

var ErrDbDataNotFound = errors.New("找不到数据")
var ErrDbInvalidArgument = errors.New("参数错误")

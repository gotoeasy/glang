package cmn

// 打印Debug级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Debug(v ...any) {
	_glc.Debug(v...)
}

// 打印Info级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Info(v ...any) {
	_glc.Info(v...)
}

// 打印Warn级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Warn(v ...any) {
	_glc.Warn(v...)
}

// 打印Error级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Error(v ...any) {
	_glc.Error(v...)
}

// 停止接收新的日志并等待日志全部输出完成
func WaitGlcFinish() {
	_glc.WaitFinish()
}

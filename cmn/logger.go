package cmn

// 打印Debug级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Debug(v ...any) {
	glc.Debug(v...)
}

// 打印Info级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Info(v ...any) {
	glc.Info(v...)
}

// 打印Warn级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Warn(v ...any) {
	glc.Warn(v...)
}

// 打印Error级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Error(v ...any) {
	glc.Error(v...)
}

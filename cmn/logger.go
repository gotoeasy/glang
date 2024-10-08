package cmn

import "runtime"

// 打印Debug级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Debug(v ...any) {
	if _glc != nil && (_glc.logLevel > 1 || (!_glc.opt.enableConsoleLog && !_glc.opt.enable)) {
		return // 关闭日志输出时，跳过
	}
	if _glc != nil && _glc.opt.PrintSrcLine {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			v = append(v, "\n  caller:"+file+":"+IntToString(line))
		}
	}
	_glc.Debug(v...)
}

// 打印Info级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Info(v ...any) {
	if _glc != nil && (_glc.logLevel > 2 || (!_glc.opt.enableConsoleLog && !_glc.opt.enable)) {
		return // 关闭日志输出时，跳过
	}
	if _glc != nil && _glc.opt.PrintSrcLine {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			v = append(v, "\n  caller:"+file+":"+IntToString(line))
		}
	}
	_glc.Info(v...)
}

// 打印Warn级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Warn(v ...any) {
	if _glc != nil && (_glc.logLevel > 3 || (!_glc.opt.enableConsoleLog && !_glc.opt.enable)) {
		return // 关闭日志输出时，跳过
	}
	if _glc != nil && _glc.opt.PrintSrcLine {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			v = append(v, "\n  caller:"+file+":"+IntToString(line))
		}
	}
	_glc.Warn(v...)
}

// 打印Error级别日志，参数将忽略nil，参数含多个GlcData时仅最后一个有效
func Error(v ...any) {
	if _glc != nil && (_glc.logLevel > 4 || (!_glc.opt.enableConsoleLog && !_glc.opt.enable)) {
		return // 关闭日志输出时，跳过
	}
	if _glc != nil && _glc.opt.PrintSrcLine {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			v = append(v, "\n  caller:"+file+":"+IntToString(line))
		}
	}
	_glc.Error(v...)
}

// 停止接收新的日志并等待日志全部输出完成
func WaitGlcFinish() {
	if _glc != nil {
		_glc.WaitFinish()
	}
}

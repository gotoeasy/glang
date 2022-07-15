package genv

import (
	"os"
	"runtime"
)

func IsWin() bool {
	return runtime.GOOS == "windows"
}

func IsMac() bool {
	return runtime.GOOS == "darwin"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func PathSeparator() string {
	return string(os.PathSeparator)
}

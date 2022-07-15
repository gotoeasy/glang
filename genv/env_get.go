package ghttp

import (
	"os"

	"github.com/gotoeasy/glang/gconv"
)

func GetString(name string, defaults ...string) string {
	s := os.Getenv(name)
	if s == "" && len(defaults) > 0 {
		return defaults[0]
	}
	return s
}

func GetInt(name string, defaults ...int) int {
	return gconv.ToInt(os.Getenv(name), defaults...)
}

func GetBool(name string, defaults ...bool) bool {
	return gconv.ToBool(os.Getenv(name), defaults...)
}

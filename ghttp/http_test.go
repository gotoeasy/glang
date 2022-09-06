package ghttp

import (
	"log"
	"testing"

	"github.com/gotoeasy/glang/gstring"
)

func Test_tostring(t *testing.T) {
	by, err := GetJson("http://baidu.com")
	log.Println(gstring.ToString(by), err)

}

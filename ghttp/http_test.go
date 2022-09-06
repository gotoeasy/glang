package ghttp

import (
	"log"
	"testing"
)

func Test_tostring(t *testing.T) {
	s, err := GetJson("http://baidu.com")
	log.Println(s, err)

}

package cmn

import (
	"testing"
)

func Test_ip_city(t *testing.T) {
	rs := GetCityByIp("127.0.0.1")
	Info(rs)
}

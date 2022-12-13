package cmn

import (
	"testing"
)

func Test_ip_city(t *testing.T) {
	rs := GetCityByIp("121.43.17.28")
	Info(rs)
}

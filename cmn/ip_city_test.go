package cmn

import (
	"testing"
)

func Test_ip_city(t *testing.T) {
	rs := GetCityByIp("127.0.0.1")
	Info(rs)
}

func Test_ip_city_by_ip9(t *testing.T) {
	rs := GetCityByIp_ip9("127.0.0.1")
	Info(rs)
}

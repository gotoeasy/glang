package cmn

import (
	"testing"
)

func Test_ip_city(t *testing.T) {
	rs := GetCityByIp("121.40.12.65")
	Info(rs)
}

package cmn

import (
	"testing"
	"time"
)

func Test_event_bus(t *testing.T) {

	e := NewEventBus()
	e.On("合计", _test_fn_sum)
	e.At("合计", 1, 2, 3)

	e.Off("合计", _test_fn_sum)
	e.At("合计", 1, 2, 3, 9999999999)

	e.On("合计", _test_fn_sum)
	e.On("合计", _test_fn_sum2)
	e.On("合计", _test_fn_sum2)

	e = NewEventBus()

	e.At("合计", 1, 2, 3, 321.4)
	e.At("合计", 1, 2, 3, 3214)
	time.Sleep(time.Second * 1)
}

func _test_fn_sum(params ...any) {
	sum := 1000
	for i := 0; i < len(params); i++ {
		sum += params[i].(int)
	}
	Info(sum)
}

func _test_fn_sum2(params ...any) {
	sum := 0
	for i := 0; i < len(params); i++ {
		sum += params[i].(int)
	}
	Info(sum)
}

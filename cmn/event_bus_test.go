package cmn

import (
	"testing"
	"time"
)

func Test_event_bus(t *testing.T) {
	e := NewEventBus()
	e.On("合计", func(params ...any) {
		sum := 0
		for i := 0; i < len(params); i++ {
			sum += params[i].(int)
		}
		Info(sum)
	})
	e.On("合计", func(params ...any) {
		sum := 1000
		for i := 0; i < len(params); i++ {
			sum += params[i].(int)
		}
		Info(sum)
	})

	e = NewEventBus()

	e.At("合计", 1, 2, 3, 3214)
	e.At("合计", 1, 2, 3, 321.4)
	e.At("合计", 1, 2, 3, 3214)
	time.Sleep(time.Second * 1)
}

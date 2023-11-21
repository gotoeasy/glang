package cmn

import (
	"testing"
)

func Test_SyncExecutor(t *testing.T) {
	s := NewSyncExecutor()
	rs := s.Exec(func(args ...any) any {
		Debug(args...)
		return args[0]
	}, 1, 2, 3, 4, 5)
	Info(rs)
}

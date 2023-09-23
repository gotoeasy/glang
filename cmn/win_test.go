package cmn

import (
	"testing"
)

func Test_win(t *testing.T) {
	Info(RegistrySetUrlProtocol("zzzz", "aaaaaaa", "C:\\Windows\\System32\\cmd.exe"))
}

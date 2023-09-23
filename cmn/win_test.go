package cmn

import (
	"testing"
)

func Test_win(t *testing.T) {
	Info(RegistrySetUrlProtocol("zzzz", "aaaaaaa", "C:\\Windows\\System32\\cmd.exe"))
}

func Test_win2(t *testing.T) {
	Debug(HostsAddDomain("127.0.0.1", "test.localhost") == nil)
}

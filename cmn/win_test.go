package cmn

import (
	"testing"
)

func Test_win(t *testing.T) {
	Info(CreateRegistry4UrlProtocol("zzzz", "ssssssssssssss", "C:\\Windows\\notepad.exe"))
}

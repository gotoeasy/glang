package cmn

import (
	"testing"

	"golang.org/x/sys/windows/registry"
)

func Test_win(t *testing.T) {
	Info(RegistrySetUrlProtocol("zzzz", "aaaaaaa", "C:\\Windows\\System32\\cmd.exe"))
}

func Test_win2(t *testing.T) {
	Debug(HostsAddDomain("127.0.0.1", "test.localhost") == nil)
}

func Test_win3(t *testing.T) {
	Debug(RegistrySetStringValue(registry.CLASSES_ROOT, "zzzz", "InstallDir", "C:\\Windows\\System32"))
	Debug(RegistryGetStringValue(registry.CLASSES_ROOT, "zzzz", "InstallDir", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

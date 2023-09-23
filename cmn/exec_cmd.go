package cmn

import (
	"bytes"
	"os/exec"
)

// 执行命令（Windows时为cmd，否则是bash）
func ExecCmd(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errout bytes.Buffer

	var cmd *exec.Cmd
	if IsWin() {
		cmd = exec.Command("cmd")
	} else {
		cmd = exec.Command("/bin/bash", "-c", command)
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err = cmd.Run()

	if err != nil {
		stderr = BytesToString(errout.Bytes())
	}
	stdout = BytesToString(out.Bytes())

	return
}

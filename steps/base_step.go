package steps

import (
	"strings"
	"os/exec"
)

func ExecSystemCommand(cmd string, shell bool) ([]byte, error) {
	var out []byte
	if shell {
		execOut, err := exec.Command("bash", "-c", cmd).CombinedOutput()
		if err != nil {
			return execOut, err
		}
		out = execOut
	} else {
		parts := strings.Fields(cmd)
		head := parts[0]
		parts = parts[1:len(parts)]

		execOut, err := exec.Command(head, parts...).CombinedOutput()
		if err != nil {
			return execOut, err
		}
		out = execOut
	}

	return out, nil
}

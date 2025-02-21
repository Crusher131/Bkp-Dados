package cmd

import (
	"os/exec"
	"syscall"

	"github.com/Crusher131/logger"
)

func CmdExec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	output, err := cmd.CombinedOutput()
	if exitError, ok := err.(*exec.ExitError); ok {
		if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
			if status.ExitStatus() != 1 {
				return err
			}
		}

	}

	logger.Info(string(output))
	return nil
}

package runner

import (
	"os/exec"
	"strings"

	logger "github.com/dterbah/gopm/log"
)

/*
Run a specific command
*/
func RunScript(command string) error {
	parts := strings.Fields(command)
	cmdName := parts[0]
	cmdArgs := parts[1:]

	cmd := exec.Command(cmdName, cmdArgs...)

	output, err := cmd.Output()

	if err != nil {
		return err
	}
	if len(output) > 0 {
		logger.Info("%s", output)
	}

	return nil
}

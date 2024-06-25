package runner

import (
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
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

	logrus.Infof("%s", output)

	return nil
}

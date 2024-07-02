package runner

import (
	"fmt"
	"os/exec"
	"strings"

	logger "github.com/dterbah/gopm/log"
)

/*
Run a specific command
*/
func RunScript(command string) error {
	var errbuf strings.Builder
	parts := strings.Fields(command)
	cmdName := parts[0]
	cmdArgs := parts[1:]

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stderr = &errbuf
	output, err := cmd.Output()

	if err != nil {
		return fmt.Errorf("error when running the command %s with the following message : %s", command, errbuf.String())
	}
	if len(output) > 0 {
		logger.Info("%s", output)
	}

	return nil
}

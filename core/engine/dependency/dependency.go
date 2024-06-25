package dependency

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/dterbah/gopm/core"
	"github.com/sirupsen/logrus"
)

/*
* Install dependencies for the current golang project. If the dependencies arg is empty, this
function will try ot install all the existing dependencies in the gopm.json
*/
func Install(dependencies []string) error {
	config, err := core.ReadConfig()

	if err != nil {
		return err
	}

	if len(dependencies) == 0 {
		for dependency, version := range config.Dependencies {
			formattedDependency := fmt.Sprintf("%s@%s", dependency, version)
			err := installDependency(formattedDependency)
			if err != nil {
				return err
			}
		}

		return nil
	} else {
		// Add dependencies to the current config, and write it in
		// the gopm.json
		for _, dependency := range dependencies {
			// find version of installed lib
			version := "latest"
			parts := strings.Split(dependency, "@")
			if len(parts) == 2 {
				version = parts[1]
			}

			err := installDependency(dependency)
			if err != nil {
				return err
			}

			config.Dependencies[parts[0]] = version
		}
	}

	return core.ExportConfig(*config, core.GOPM_CONFIG_FILE)
}

// ---- Private functions ---- //

/*
Install specific dependency in the project
*/
func installDependency(dependency string) error {
	logrus.Infof("⏳ Instaling dependency %s ...", dependency)
	cmd := exec.Command("go", "get", dependency)
	// todo: error case
	err := cmd.Run()

	if err != nil {
		return err
	}

	logrus.Infof("✅ Dependency %s installed !", dependency)
	return nil
}

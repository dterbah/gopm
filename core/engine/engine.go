package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dterbah/gopm/core/config"
	"github.com/dterbah/gopm/core/license"
	"github.com/dterbah/gopm/utils/file"
	"github.com/sirupsen/logrus"
)

/*
Name of the gopm file used by the CLI
*/
const GOPM_CONFIG_FILE = "gopm.json"
const LICENSE_FILE = "LICENSE.txt"

const DEFAULT_ENTRY_POINT_CONTENT = `
package main

import "fmt"

func main() {
	fmt.Println("Hello world !")
}
`

/*
Init a project with given user information
*/
func InitProject(config config.GoPMConfig) error {
	// setup the basic scripts in the configurations
	config.Scripts["build"] = "go build"
	config.Scripts["run"] = "go run " + config.EntryPoint
	config.Scripts["test"] = "go test ./..."

	/*
		init steps : create the directory associated with the project name,
		then create the gopm.json, entry point, fetch license,
		and execute go mod init
	*/
	steps := []struct {
		name string
		fn   func() error
	}{
		{"create project directory", func() error { return createProjectDirectory(config.ProjectName) }},
		{"export config", func() error { return exportConfig(config, filepath.Join(config.ProjectName, GOPM_CONFIG_FILE)) }},
		{"create entry point", func() error { return createEntryPoint(config.EntryPoint, config.ProjectName) }},
		{"fetch and export license", func() error {
			licenseContent, err := license.FetchLicense(config.License)
			if err != nil {
				return err
			}
			return exportLicense(config.ProjectName, licenseContent)
		}},
		{"initialize Go project", func() error { return initGoProject(config.ProjectName, config.Git) }},
	}

	for _, step := range steps {
		if err := step.fn(); err != nil {
			return fmt.Errorf("failed to %s: %w", step.name, err)
		}
	}

	return nil
}

/*
* Add dependencies for the current golang project
 */
func AddDependencies(dependencies []string) error {
	config, err := config.ReadConfig(GOPM_CONFIG_FILE)

	if err != nil {
		return err
	}

	// Add dependencies to the current config, and write it in
	// the gopm.json
	for _, dependency := range dependencies {
		// find version of installed lib
		version := "latest"
		parts := strings.Split(dependency, "@")
		if len(parts) == 2 {
			version = parts[1]
		}

		logrus.Infof("⏳ Instaling dependency %s ...", dependency)
		cmd := exec.Command("go", "get", dependency)
		// todo: error case
		err := cmd.Run()

		if err != nil {
			return err
		}

		config.Dependencies[dependency] = version
		logrus.Infof("✅ Dependency %s installed !", dependency)
	}

	return exportConfig(*config, GOPM_CONFIG_FILE)
}

// ---- Private functions ---- //

/*
Create the project directory
*/
func createProjectDirectory(dirName string) error {
	return os.Mkdir(dirName, os.ModePerm)
}

func createEntryPoint(entryPoint string, toDir string) error {
	entryPointPath := filepath.Join(toDir, entryPoint)

	file, err := os.OpenFile(entryPointPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("error when creating entry point")
	}

	_, err = file.WriteString(DEFAULT_ENTRY_POINT_CONTENT)
	if err != nil {
		return errors.New("erorr when writing the entry point")
	}

	err = file.Close()

	return err
}

/*
Export the license in a file
*/
func exportLicense(projectName, licenseContent string) error {
	licensePath := filepath.Join(projectName, LICENSE_FILE)
	file, err := os.OpenFile(licensePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New("license file failed")
	}

	defer file.Close()

	_, err = file.Write([]byte(licenseContent))
	if err != nil {
		return errors.New("writing license file")
	}

	return nil
}

func initGoProject(projectDir string, repositoryName string) error {
	dir, _ := os.Getwd()
	cmd := exec.Command("go", "mod", "init", repositoryName)
	cmd.Dir = filepath.Join(dir, projectDir)
	return cmd.Run()
}

/*
Export the configuration in a file
*/
func exportConfig(config config.GoPMConfig, configPath string) error {
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New("config to json failed")
	}

	flags := os.O_WRONLY | os.O_CREATE
	if !file.IsExists(configPath) {
		flags = flags | os.O_APPEND
	}

	file, err := os.OpenFile(configPath, flags, 0644)
	if err != nil {
		return errors.New("gopm config file failed")
	}

	defer file.Close()

	_, err = file.Write(configJSON)
	if err != nil {
		return errors.New("writing gopm config file")
	}

	return nil
}

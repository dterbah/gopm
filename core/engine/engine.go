package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dterbah/gopm/core/config"
	"github.com/dterbah/gopm/core/license"
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
		{"export config", func() error { return exportConfig(config) }},
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
func exportConfig(config config.GoPMConfig) error {
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New("config to json failed")
	}

	configPath := filepath.Join(config.ProjectName, GOPM_CONFIG_FILE)

	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

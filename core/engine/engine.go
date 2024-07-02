package engine

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dterbah/gopm/core"
	"github.com/dterbah/gopm/core/engine/dependency"
	runner "github.com/dterbah/gopm/core/engine/script-runner"
)

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
func InitProject(config core.GoPMConfig) error {
	// setup the basic scripts in the configurations
	config.Scripts["build"] = "go build"
	config.Scripts["main"] = "go run " + config.EntryPoint
	config.Scripts["test"] = "go test ./..."
	config.Scripts["fmt"] = "go fmt ./..."
	config.Scripts["tidy"] = "go mod tidy"
	config.Scripts["clean"] = "go clean"

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
		{"export config", func() error {
			return core.ExportConfig(config, filepath.Join(config.ProjectName, core.GOPM_CONFIG_FILE))
		}},
		{"create entry point", func() error { return createEntryPoint(config.EntryPoint, config.ProjectName) }},
		{"fetch and export license", func() error {
			licenseContent, err := core.FetchLicense(config.License)
			if err != nil {
				return err
			}
			return core.ExportLicense(config.ProjectName, licenseContent)
		}},
		{"initialize Go project", func() error { return initGoProject(config.ProjectName, config.Git) }},
	}

	for _, step := range steps {
		if err := step.fn(); err != nil {
			os.RemoveAll(config.ProjectName)
			return fmt.Errorf("failed to %s: %w", step.name, err)
		}
	}

	return nil
}

/*
* Install dependencies for the current golang project. If the dependencies arg is empty, this
function will try ot install all the existing dependencies in the gopm.json
*/
func InstallDependencies(dependencies []string) error {
	return dependency.Install(dependencies)
}

func RunScript(scriptName string) error {
	config, err := core.ReadConfig()

	if err != nil {
		return err
	}

	// check if command exists
	scriptCommand, ok := config.Scripts[scriptName]
	if !ok {
		return fmt.Errorf("script %s not found", scriptName)
	}

	return runner.RunScript(scriptCommand)
}

// ---- Private functions ---- //

/*
Create the project directory
*/
func createProjectDirectory(dirName string) error {
	if len(dirName) == 0 {
		return errors.New("dir name empty")
	}
	return os.Mkdir(dirName, os.ModePerm)
}

func createEntryPoint(entryPoint string, toDir string) error {
	if len(entryPoint) == 0 {
		return errors.New("empty entry point")
	}
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

func initGoProject(projectDir string, repositoryName string) error {
	dir, _ := os.Getwd()
	cmd := exec.Command("go", "mod", "init", repositoryName)
	cmd.Dir = filepath.Join(dir, projectDir)
	return cmd.Run()
}

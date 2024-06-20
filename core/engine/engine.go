package engine

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/dterbah/gopm/core/config"
)

/*
Name of the gopm file used by the CLI
*/
const GOPM_CONFIG_FILE = "gopm.json"

/*
Init a project with given user information
*/
func InitProject(config config.GoPMConfig) error {
	return exportConfig(config)
}

/*
Export the configuration in a file
*/
func exportConfig(config config.GoPMConfig) error {
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return errors.New("config to json failed")
	}

	// Écrire le JSON dans un fichier
	file, err := os.Create(GOPM_CONFIG_FILE)
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

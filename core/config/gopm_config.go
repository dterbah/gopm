package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

/*
Defines the content of the gopm.json. This config will be used
to store the different information about the current go project
*/
type GoPMConfig struct {
	Author       string            `json:"author"`       // Author of the project
	Description  string            `json:"description"`  // Description of the project
	Version      string            `json:"version"`      // version of the project
	EntryPoint   string            `json:"entry"`        // entry point of the project (e.g main.go)
	License      string            `json:"license"`      // Licence of the project
	Scripts      map[string]string `json:"scripts"`      // Map of scripts that can be used
	ProjectName  string            `json:"name"`         // Name of the project
	Git          string            `json:"git"`          // Git repository name
	Dependencies map[string]string `json:"dependencies"` // Dependencies of the project
}

/*
Create an empty configuration
*/
func NewGoPMConfig() *GoPMConfig {
	return &GoPMConfig{
		Scripts:      make(map[string]string),
		Dependencies: make(map[string]string),
	}
}

/*
Read the gopm configuration from a file
*/
func ReadConfig(path string) (*GoPMConfig, error) {
	config := NewGoPMConfig()
	configFile, err := os.Open(path)

	if err != nil {
		return nil, errors.New("error when opening the configuration file")
	}

	defer configFile.Close()

	byteValue, err := io.ReadAll(configFile)
	if err != nil {
		return nil, errors.New("error when reading the configuration file")
	}

	err = json.Unmarshal(byteValue, config)
	if err != nil {
		logrus.Error(err)
		return nil, errors.New("error when parsing the configuration file")
	}

	return config, nil
}

package config

/*
Defines the content of the gopm.json. This config will be used
to store the different information about the current go project
*/
type GoPMConfig struct {
	Author      string            `json:"author"`      // Author of the project
	Description string            `json:"description"` // Description of the project
	Version     string            `json:"version"`     // version of the project
	EntryPoint  string            `json:"entry"`       // entry point of the project (e.g main.go)
	License     string            `json:"license"`     // Licence of the project
	Scripts     map[string]string `json:"scripts"`     // Map of scripts that can be used
	ProjectName string            `json:"name"`        // Name of the project
	Git         string            `json:"name"`        // Git repository name
}

/*
Create an empty configuration
*/
func NewGoPMConfig() *GoPMConfig {
	return &GoPMConfig{
		Scripts: make(map[string]string),
	}
}

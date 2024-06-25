package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const GITHUB_LICENSE_BASE_URL = "https://api.github.com/licenses"
const LICENSE_FILE = "LICENSE.txt"

/*
Response JSON from the Github API for retrieving licence templates
*/
type License struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Body string `json:"body"`
}

func FetchLicense(name string) (string, error) {
	url := fmt.Sprintf("%s/%s", GITHUB_LICENSE_BASE_URL, name)

	response, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch license: status code %d", response.StatusCode)
	}

	var license License
	err = json.NewDecoder(response.Body).Decode(&license)
	if err != nil {
		return "", err
	}

	return license.Body, nil
}

/*
Export the license in a file
*/
func ExportLicense(projectName, licenseContent string) error {
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

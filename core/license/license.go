package license

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const GITHUB_LICENSE_BASE_URL = "https://api.github.com/licenses"

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

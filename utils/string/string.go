package strutil

import "strings"

/*
If the input is empty, this function wil return the default value.
Else, the input value is returned
*/
func GetStringIfEmpty(input string, defaultValue string) string {
	cleanInput := strings.TrimSpace(input)
	if cleanInput == "" {
		return defaultValue
	}
	return cleanInput
}

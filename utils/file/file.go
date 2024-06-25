package file

import (
	"os"
)

/*
Check if file exists or not
*/
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

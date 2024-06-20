package user

import (
	"os/user"
)

/*
Return the username of the current session
*/
func GetUsername() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}

	return user.Username
}

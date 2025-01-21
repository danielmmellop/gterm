package os

import "os/user"

func GetUserName() string {
	user, _ := user.Current()
	return user.Username
}

func GetHomeDir() string {
	user, _ := user.Current()
	return user.HomeDir
}

package os

var CurrentPath string

func SetCurrentFolder() string {
	CurrentPath = GetHomeDir()
	return CurrentPath
}

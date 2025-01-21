package test

import "os"

func GetTestDir() string {
	return os.TempDir()
}

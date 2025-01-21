package os

import (
	"fmt"
	"os"
)

func Env(params []string) ([]string, error) {
	var results []string
	var err error
	listEnvName := "list"

	switch true {
	case len(params) == 1:
		if params[0] != listEnvName {
			value := get(params[0])
			results = append(results, value)
		} else {
			results = getAll()
		}
	case len(params) == 2:
		if params[0] == listEnvName {
			return []string{}, fmt.Errorf("not allowed to create a env called test")
		}
		err = set(params[0], params[1])
	}

	return results, err
}

func set(name string, value string) error {
	return os.Setenv(name, value)
}

func get(name string) string {
	return os.Getenv(name)
}

func getAll() []string {
	return os.Environ()
}

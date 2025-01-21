package directory

import (
	"errors"
	"fmt"
	"gterm/internal/os"
	goos "os"
	"path/filepath"
)

func Change(params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("directory not found")
	}

	path, err := filepath.Abs(params[0])
	if err != nil {
		return errors.New("directory not found")
	}
	err = goos.Chdir(path)
	if err != nil {
		return errors.New("directory not found")
	}

	os.CurrentPath = path

	return nil
}

func Create(params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("missing operand")
	}
	path := params[0]
	fullPath, _ := filepath.Abs(path)
	err := goos.MkdirAll(fullPath, goos.ModePerm)
	if err != nil && !goos.IsExist(err) {
		return errors.New("create " + path + " fail")
	}

	return nil
}

func Delete(params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("cannot remove : No such file or directory")
	}

	path := params[0]
	fullPath, _ := filepath.Abs(path)
	err := goos.RemoveAll(fullPath)
	if err != nil && !goos.IsExist(err) {
		return errors.New("delete " + path + " fail")
	}

	return nil
}

func List(params []string) ([]string, error) {
	var path string

	if len(params) == 0 {
		path = os.CurrentPath
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return []string{}, err
	}

	entries, err := goos.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, e := range entries {
		files = append(files, e.Name())
	}

	return files, nil
}

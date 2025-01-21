package file

import (
	"bufio"
	"fmt"
	"gterm/internal/directory"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Create(params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("cannot create")
	}

	path := params[0]
	fullPath, _ := filepath.Abs(path)
	if isFolderPath(fullPath) {
		return fmt.Errorf("cannot create '%s'", path)
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func Remove(path string) error {
	return os.RemoveAll(path)
}

func Read(params []string) ([]string, error) {
	if len(params) != 1 {
		return []string{}, fmt.Errorf("no such file")
	}

	file, err := os.Open(params[0])
	if err != nil {
		return []string{}, err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.TrimSpace(scanner.Text()))
	}

	return data, nil
}

// Source is a File
// Target is a Folder
func Move(params []string, deleteSource bool) error {
	switch true {
	case len(params) < 2:
		return fmt.Errorf("missing file operand")
	case len(params) == 1:
		return fmt.Errorf("missing destination")
	}
	source := params[0]
	isSourceFilePath := IsFilePath(source)
	if !isSourceFilePath {
		return fmt.Errorf("cannot move '%s': Is a directory", source)
	}

	target := params[1]
	isTargetFolderPath := isFolderPath(target)
	if !isTargetFolderPath {
		return fmt.Errorf("cannot move '%s': %s is a file", source, target)
	}

	fullSourcePath, err := filepath.Abs(source)
	if err != nil {
		return fmt.Errorf("cannot move '%s'", source)
	}

	fullTargetPath, _ := filepath.Abs(target)
	fileTargetName := filepath.Base(source)
	completeTargetPath := fmt.Sprintf("%s/%s", fullTargetPath, fileTargetName)
	_, err = os.Stat(fullTargetPath)
	if os.IsNotExist(err) {
		directory.Create([]string{fullTargetPath})
	}

	// fullSourcePath should point to a file
	// fullTargetPath should point to a file
	errCopy := Copy(fullSourcePath, completeTargetPath)
	if errCopy != nil {
		return errCopy
	}

	if deleteSource {
		errRemove := Remove(fullSourcePath)
		if errRemove != nil {
			return errRemove
		}
	}

	return nil
}

func IsFilePath(path string) bool {
	sourceExt := filepath.Ext(path)
	return len(sourceExt) != 0
}

func isFolderPath(path string) bool {
	sourceExt := filepath.Ext(path)
	return len(sourceExt) == 0
}

func Copy(source string, target string) error {
	r, err := os.Open(source)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(target)
	if err != nil {
		return err
	}
	defer w.Close()
	w.ReadFrom(r)

	return nil
}

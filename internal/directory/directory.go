package directory

import (
	"io/ioutil"
	"os"
	"strings"
)

// List returns a slice of directory names in the given folder
func List(folderPath string) ([]string, error) {
	var directories []string

	items, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.IsDir() {
			directories = append(directories, item.Name())
		}
	}

	return directories, nil
}

// FindTypeScriptFiles returns a slice of TypeScript (.ts) file names in the given folder
func FindTypeScriptFiles(folderPath string) ([]string, error) {
	var tsFiles []string

	items, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if !item.IsDir() && strings.HasSuffix(item.Name(), ".ts") {
			tsFiles = append(tsFiles, item.Name())
		}
	}

	return tsFiles, nil
}

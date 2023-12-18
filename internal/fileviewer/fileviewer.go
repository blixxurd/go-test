package fileviewer

import (
	"os"
)

// DisplayFileContents reads and returns the contents of the given file as a string
func DisplayFileContents(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

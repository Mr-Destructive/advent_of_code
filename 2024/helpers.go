package helpers

import (
	"bytes"
	"os"
)

func ReadFileBytes(path string) []byte {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return fileBytes
}

func ReadFileLines(path string) []string {
	fileBytes := ReadFileBytes(path)
	lines := []string{}
	separator := []byte("\n")
	for _, line := range bytes.Split(fileBytes, separator) {
		if string(line) != "" {
			lines = append(lines, string(line))
		}
	}
	return lines
}

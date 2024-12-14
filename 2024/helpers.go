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

func ReadFileSections(path string) [][]string {
	fileBytes := ReadFileBytes(path)
	lines := [][]string{}
	separator := []byte("\n\n")
	for _, section := range bytes.Split(fileBytes, separator) {
		str := []string{}
		if string(section) != "" {
			for _, line := range bytes.Split(section, []byte("\n")) {
				str = append(str, string(line))
			}
			lines = append(lines, str)
		}
	}
	return lines
}

package main

import (
	"os"
	"strings"
)

func ReadInput(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	file_content := string(content)
	data := strings.Split(file_content, "\n")
	data = data[:len(data)-1]
	return data, err
}

func main() {
	lines, err := ReadInput("../inputs/test/d04.txt")
	if err != nil {
		panic(err)
	}
	for  _ = range lines {
	}
}

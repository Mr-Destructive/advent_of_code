package helpers

import (
	"log"
	"os"
	"strings"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func ReadInput(file_path string) ([]string, error) {
	file, err := os.ReadFile(file_path)
	HandleError(err)
	file_content := strings.Split(string(file), "\n")
	file_content = file_content[:len(file_content)-1]
	return file_content, err
}

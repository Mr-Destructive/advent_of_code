package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	file_content := string(content)
	data := strings.Split(file_content, "\n")
	return data, err
}

func replaceWordsToDigits(line string) string {
	words := map[string]string{
		"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}
	for word := range words {
		lastPos := strings.LastIndex(line, word)
		if lastPos != -1 {
			line = line[:lastPos] + words[word] + line[lastPos+len(word):]
		}
	}
	return line
}

func getDigits(line string) int {
	str := ""
	n := len(line)

	for i := 0; i < n; i++ {
		ch := line[i]
		if ch >= '0' && ch <= '9' {
			str += string(ch)
		} else if len(str) > 0 {
			break
		}
	}
	res, _ := strconv.Atoi(str)
	return res
}

func main() {
	data, err := ReadInput("../inputs/test/d01-2.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, line := range data {
		sum += getDigits(line)
	}
	fmt.Println(sum)

}

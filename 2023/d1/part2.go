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
	return ""

}

func getDigits(line string) int {
	str := ""
	n := len(line)
	for _, ch := range line {
		if ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9' {
			str += string(ch)
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		ch := line[i]
		if ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' || ch == '7' || ch == '8' || ch == '9' {
			str += string(ch)
			break
		}
	}
	fmt.Println(str)

	res, _ := strconv.Atoi(str)
	return res
}

func main() {
	data, err := ReadInput("part1.test")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, line := range data {
		sum += getDigits(line)
	}
	fmt.Println(sum)

}

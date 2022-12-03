package main

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

func CharToPoints(char byte) int {
	char_point := char
	if char_point > 64 && char_point < 92 {
		char_point -= 38
	} else {
		char_point -= 96
	}
	return int(char_point)
}

func GetCommonChars(s1, s2 string) string {
	common_chars := ""
	for _, s := range s1 {
		for _, p := range s2 {
			if s == p && !strings.Contains(common_chars, string(s)) {
				common_chars += string(s)
			}
		}
	}
	return common_chars
}

func main() {
	file, err := os.ReadFile("input3.txt")
	HandleError(err)
	file_content := strings.Split(string(file), "\n")
	file_content = file_content[:len(file_content)-1]

	total_points := 0
	for _, compartment := range file_content {
		length := len(compartment) / 2
		left, right := compartment[:length], compartment[length:]
		common_char := strings.IndexAny(left, right)
		char_point := CharToPoints(compartment[common_char])
		total_points += char_point
	}
	log.Println(total_points)
	// PART 1: 7746

	length := len(file_content)
	total_points = 0
	for i := 0; i < length; i += 3 {
		common_chars := GetCommonChars(file_content[i], file_content[i+1])
		common_chars_1 := GetCommonChars(file_content[i+1], file_content[i+2])
		var char_point int
		for _, s := range common_chars {
			common_char := strings.Index(common_chars_1, string(s))
			if common_char != -1 {
				char_point = CharToPoints(common_chars_1[common_char])
				break
			}
		}
		total_points += char_point
	}
	log.Println(total_points)
	// PART 2: 2604
}

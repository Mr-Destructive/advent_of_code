package main

import (
	"log"
	"strings"

	"aoc2022/helpers"
)

func CharToPoints(char byte) int {
	char_point := char
	if char_point > 64 && char_point < 92 {
		char_point -= 38
	} else {
		char_point -= 96
	}
	return int(char_point)
}

func main() {
	file_content, err := helpers.ReadInput("inputs/input3.test")
	helpers.HandleError(err)

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
		for _, s := range file_content[i] {
			if strings.ContainsRune(file_content[i+1], s) && strings.ContainsRune(file_content[i+2], s) {
				total_points += CharToPoints(byte(s))
				break
			}
		}
	}
	log.Println(total_points)
	// PART 2: 2604
}

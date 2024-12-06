package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFileBytes(path string) []byte {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
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

func Multiply(line string) int {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := mulRegex.FindAllStringSubmatch(line, -1)
	score := 0
	for _, match := range matches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		score += x * y
	}
	return score
}

func StripDoDont(line string) string {
	enabled := true
	result := ""
	dontOffset := len("don't()")
	doOffset := len("do()")

	for len(line) > 0 {
		dontIndex := strings.Index(line, "don't()")
		doIndex := strings.Index(line, "do()")

		if dontIndex == -1 && doIndex == -1 {
			if enabled {
				result += line
			}
			break
		}
		if dontIndex != -1 && (dontIndex < doIndex || doIndex == -1) {
			if enabled {
				result += line[:dontIndex]
			}
			enabled = false
			line = line[dontIndex+dontOffset:]
		} else {
			if enabled {
				result += line[:doIndex]
			}
			enabled = true
			line = line[doIndex+doOffset:]
		}
	}
	return result
}

func main() {
	lines := ReadFileLines("../../inputs/day03/prod.txt")
	score := 0
	//for _, line := range lines {
	//	score += Multiply(line)
	//}
	//fmt.Println(score)
	score = 0
	lineStr := ""
	for _, line := range lines {
		lineStr += line
	}
	score += Multiply(StripDoDont(lineStr))
	fmt.Println(score)
}

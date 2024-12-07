package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func GetValEquation(lines []string) [][]int {
	nums := [][]int{}
	for _, lineStr := range lines {
		numLine := []int{}
		for i, numStr := range strings.Split(lineStr, " ") {
			if i == 0 {
				numStr = strings.Trim(numStr, ":")
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			numLine = append(numLine, num)
		}
		nums = append(nums, numLine)
	}
	return nums
}

func IsCalibrated(expected, value int, equations []int) bool {
	if len(equations) == 0 {
		return expected == value
	}
	if value > expected {
		return false
	}
    val := value+equations[0]
	if IsCalibrated(expected, val, equations[1:]) {
		return true
	}
    val = value * equations[0]
	return IsCalibrated(expected, val, equations[1:])
}

func main() {
	lines := ReadFileLines("../../inputs/day07/prod.txt")
	nums := GetValEquation(lines)
	score := 0
	for _, numLine := range nums {
		if IsCalibrated(numLine[0], 0, numLine[1:]) {
			score += numLine[0]
		}
	}
	fmt.Println(score)
}

package main

import (
	"bytes"
	"fmt"
	"math"
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
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic(fmt.Sprintf("Invalid line format: %s", line))
		}

		expected, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic(err)
		}

		operands := strings.Fields(parts[1])
		numLine := []int{expected}
		for _, numStr := range operands {
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

func IsCalibrated(expected int, equations []int) bool {
	queue := [][]int{{1, equations[0]}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		index := current[0]
		result := current[1]

		if index == len(equations) {
			if result == expected {
				return true
			}
			continue
		}

		nextNum := equations[index]
		queue = append(queue, []int{index + 1, result + nextNum})
		queue = append(queue, []int{index + 1, result * nextNum})
	}
	return false
}

type Queue struct {
	index  int
	result int
	path   string
}

func IsCalibratedConcat(expected int, equations []int) bool {
	queue := []Queue{
		{0, equations[0], fmt.Sprintf("%d", equations[0])},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.index == len(equations)-1 {
			if current.result == expected {
				return true
			}
			continue
		}

		if current.index >= len(equations)-1 {
			continue
		}

		nextNum := equations[current.index+1]

		queue = append(queue, Queue{
			index:  current.index + 1,
			result: current.result + nextNum,
			path:   fmt.Sprintf("%s + %d", current.path, nextNum),
		})

		queue = append(queue, Queue{
			index:  current.index + 1,
			result: current.result * nextNum,
			path:   fmt.Sprintf("%s * %d", current.path, nextNum),
		})

		concatenated := current.result*int(math.Pow10(len(fmt.Sprintf("%d", nextNum)))) + nextNum
		queue = append(queue, Queue{
			index:  current.index + 1,
			result: concatenated,
			path:   fmt.Sprintf("%s || %d", current.path, nextNum),
		})
	}

	return false
}

func main() {
	lines := ReadFileLines("../../inputs/day07/sample.txt")
	nums := GetValEquation(lines)
	score1 := 0
	score2 := 0
	for _, numLine := range nums {
		if IsCalibrated(numLine[0], numLine[1:]) {
			score1 += numLine[0]
		}
		if IsCalibratedConcat(numLine[0], numLine[1:]) {
			score2 += numLine[0]
		}
	}
	fmt.Println(score1, score2)
}

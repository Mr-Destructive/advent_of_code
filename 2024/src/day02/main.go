package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
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

func SplitLevels(lines []string) [][]int {
	reportLevels := [][]int{}
	for i, reportLine := range lines {
		reportLevels = append(reportLevels, []int{})
		for _, levelStr := range strings.Split(reportLine, " ") {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				log.Fatal(err)
			}
			reportLevels[i] = append(reportLevels[i], level)
		}
	}
	return reportLevels
}

func IsSafe(report []int) (bool, int) {
	prevDiff := report[1] - report[0]
	isIcreasing := prevDiff > 0
	if prevDiff == 0 || math.Abs(float64(prevDiff)) > 3 {
		return false, 0
	}

	for i := 2; i < len(report); i++ {
		currDiff := report[i] - report[i-1]
		if isIcreasing {
			if currDiff <= 0 || currDiff > 3 {
				return false, i - 1
			}
		} else {
			if currDiff >= 0 || currDiff < -3 {
				return false, i - 1
			}
		}

	}
	return true, -1
}

func RemoveAndCheck(report []int, index int) bool {
	if index > len(report)-1 || index < 0 {
		return false
	}
	reportNew := append([]int{}, report[:index]...)
	reportNew = append(reportNew, report[index+1:]...)
	safe, _ := IsSafe(reportNew)
	fmt.Println(safe, report)
	return safe
}

func RemoveLevels(report []int) bool {
	safe, unsafeIndex := IsSafe(report)
	if safe {
		return true
	} else {
		if RemoveAndCheck(report, unsafeIndex) {
			return true
		}
		if RemoveAndCheck(report, unsafeIndex-1) {
			return true
		}
		if RemoveAndCheck(report, unsafeIndex+1) {
			return true
		}
		return false
	}
}

func main() {
	lines := ReadFileLines("../../inputs/day02/prod.txt")
	reportLevels := SplitLevels(lines)
	safeCount := 0
	for _, report := range reportLevels {
		if ok, _ := IsSafe(report); ok {
			safeCount++
		}
	}
	//fmt.Println(safeCount)
	safeCount = 0
	for _, report := range reportLevels {
		if RemoveLevels(report) {
			safeCount++
		}
	}
	fmt.Println(safeCount)

}

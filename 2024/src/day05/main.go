package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
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

func ReadFileSections(path string) [][][]int {
	fileBytes := ReadFileBytes(path)
	lines := []string{}
	separator := []byte("\n\n")
	for _, line := range bytes.Split(fileBytes, separator) {
		if string(line) != "" {
			lines = append(lines, string(line))
		}
	}
	sections := [][][]int{}
	for i, section := range lines {
		nums := [][]int{}
		lineStrs := strings.Split(section, "\n")
		separator := ","
		if i == 0 {
			separator = "|"
		}
		for _, lineStr := range lineStrs {
			if lineStr == "" {
				continue
			}
			numL := []int{}
			for _, numStr := range strings.Split(lineStr, separator) {
				num, _ := strconv.Atoi(numStr)
				numL = append(numL, num)
			}
			nums = append(nums, numL)
		}
		sections = append(sections, nums)
	}
	return sections
}

func ConstructRules(rulesList [][]int) map[int][]int {
	rules := make(map[int][]int)
	for _, rule := range rulesList {
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	return rules
}

func GetPageIndices(rules map[int][]int, pages [][]int) map[int][]int {
	nums := make(map[int]bool)
	for num, list := range rules {
		nums[num] = true
		for _, elem := range list {
			if !nums[elem] {
				nums[elem] = true
			}
		}
	}

	numIndices := make(map[int][]int)
	for num, _ := range nums {
		for _, numLine := range pages {
			index := -1
			for i, n := range numLine {
				if n == num {
					index = i
				}
			}
			numIndices[num] = append(numIndices[num], index)
		}
	}
	return numIndices
}

func GetOrderedPages(rules, numIndices map[int][]int, pages [][]int) int {

	score := 0
	//fmt.Println(rules)
	//fmt.Println(numIndices)
	for index, pageLine := range pages {
		ordered := true
		for _, num1 := range pageLine {
			rule := rules[num1]
			index1 := numIndices[num1][index]
			for _, num2 := range rule {
				index2 := numIndices[num2][index]
				if index1 == -1 || index2 == -1 {
					continue
				}
				if index1 > index2 {
					//fmt.Println(pageLine, num1, num2, index1, index2)
					ordered = false
				}
			}
		}
		if ordered {
			score += pageLine[int(len(pageLine)/2)]
		}
	}
	return score
}

func CorrectPageOrder(line []int, rules map[int][]int) []int {
	newLine := []int{}
	for _, num := range line {
		index := make(map[int]int)
		for i, n := range newLine {
			index[n] = i
		}
		newInsertIndex := len(newLine)
		for _, rule := range rules[num] {
			if idx, ok := index[rule]; ok {
				if newInsertIndex > idx {
					newInsertIndex = idx
				}
			}
		}
		afterNum := slices.Clone(newLine[newInsertIndex:])
		newLine = append(newLine[:newInsertIndex], num)
		newLine = append(newLine, afterNum...)
	}
	return newLine
}

func GetCorrectOrderedPages(rules, numIndices map[int][]int, pages [][]int) int {

	score := 0
	for index, pageLine := range pages {
		ordered := true
		for _, num1 := range pageLine {
			rule := rules[num1]
			index1 := numIndices[num1][index]
			for _, num2 := range rule {
				index2 := numIndices[num2][index]
				if index1 == -1 || index2 == -1 {
					continue
				}
				if index1 > index2 {
					ordered = false
				}
			}
		}
		if !ordered {
			newLine := CorrectPageOrder(pageLine, rules)
			score += newLine[len(newLine)/2]
		}
	}
	return score
}

func main() {
	sections := ReadFileSections("../../inputs/day05/prod.txt")
	rules := ConstructRules(sections[0])
	pages := sections[1]
	pageIndices := GetPageIndices(rules, pages)
	//score := GetOrderedPages(rules, pageIndices, pages)
	score := GetCorrectOrderedPages(rules, pageIndices, pages)
	fmt.Println(score)

}

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadInput(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	file_content := string(content)
	data := strings.Split(file_content, "\n")
	data = data[:len(data)-1]
	return data, err
}

func splitter(line string) []int {
	winners := []int{}
	parts := strings.Split(line, "|")
	parts[0] = strings.Trim(parts[0], " ")
	parts[1] = strings.Trim(parts[1], " ")
	ourCards := strings.Split(strings.Replace(strings.Split(parts[0], ":")[1], "  ", " ", -1), " ")
	winningCards := strings.Split(strings.Replace(parts[1], "  ", " ", -1), " ")

	for _, ourCard := range ourCards {
		for _, winner := range winningCards {
			if ourCard == winner {
				num, _ := strconv.Atoi(string(winner))
				winners = append(winners, num)
			}
		}
	}
	return winners
}

func main() {
	lines, err := ReadInput("../inputs/test/d04.txt")
	if err != nil {
		panic(err)
	}
	count := 0
	for i, line := range lines {
		c := splitter(line)
		fmt.Println(i+1, len(c))
		count += int(math.Pow(2, float64(len(c)-1)))
	}
	fmt.Println(count)
}

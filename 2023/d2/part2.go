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
	data = data[:len(data)-1]
	return data, err
}

func splitter(line string) ([]int, int) {
	parts := strings.Split(line, ":")
	gameIDStr := strings.Split(parts[0], " ")[1]
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		panic(err)
	}
	sets := strings.Split(parts[1], ";")
	result := []int{0, 0, 0}
	for _, set := range sets {
		cubes := strings.Split(set, ", ")
		for _, cube := range cubes {
			clean := strings.Trim(cube, " ")
			contents := strings.Split(clean, " ")
			numberStr := contents[0]
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				panic(err)
			}
			color := contents[1]
			if color == "red" && result[0] < number {
				result[0] = number
			} else if color == "green" && result[1] < number {
				result[1] = number
			} else if color == "blue" && result[2] < number {
				result[2] = number
			}
		}
	}
	return result, gameID
}

func main() {
	lines, err := ReadInput("../inputs/input/d02.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, line := range lines {
		cubes, _ := splitter(line)
		fmt.Println(cubes)
		score := cubes[0] * cubes[1] * cubes[2]
		sum += score
	}
	fmt.Println(sum)

}

package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"aoc2022/helpers"
)

func ReadInput(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	file_content := string(content)
	data := strings.Split(file_content, "\n")
	return data, err
}

func main() {

	data, err := ReadInput("inputs/input1.test")
	//data, err := ReadInput("input1.txt")
	helpers.HandleError(err)
	calories_list := []int{}
	total_calories := 0
	max_calory := total_calories
	for _, d := range data {
		if d != "" {
			calories, err := strconv.Atoi(d)
			helpers.HandleError(err)
			total_calories += calories
		} else {
			calories_list = append(calories_list, total_calories)
			if total_calories > max_calory {
				max_calory = total_calories
			}
			total_calories = 0
		}
	}
	log.Println("Max Calories: ", max_calory)
	// Part One 72718

	sum_calories := 0
	sort.Ints(calories_list)
	start := len(calories_list) - 3
	for _, calory := range calories_list[start:] {
		log.Println(calory)
		sum_calories += calory
	}
	log.Println("Sum of all Calories: ", sum_calories)
	// Part Two 213089
}

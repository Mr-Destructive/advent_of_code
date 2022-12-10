package main

import (
	"log"
	"os"
	"strings"

	"aoc2022/helpers"
)

func ReadInput(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	file_content := string(content)
	data := strings.Split(file_content, "\n")
	data = data[:len(data)-1]
	return data, err
}

func GetScore(s string) int {
	points := 0
	switch s {
	case "A", "X":
		{
			points += 1
		}
	case "B", "Y":
		{
			points += 2
		}
	case "C", "Z":
		{
			points += 3
		}
	}
	return points
}
func AssignScore(opt_ply []string) int {
	opt := GetScore(opt_ply[0])
	ply := GetScore(opt_ply[1])
	if opt == ply {
		ply += 3
	}
	//println(opt, ply)
	if opt_ply[0] == "A" && opt_ply[1] == "Y" {
		ply += 6
	} else if opt_ply[0] == "C" && opt_ply[1] == "X" {
		ply += 6
	} else if opt_ply[0] == "B" && opt_ply[1] == "Z" {
		ply += 6
	}
	return ply
}

func GetMoveScore(opt_ply []string) int {
	opt := GetScore(opt_ply[0])
	ply := GetScore(opt_ply[1])
	score := 0
	switch ply {
	case 1:
		{
			switch opt {
			case 1:
				{
					score += 3
				}
			case 2:
				{
					score += 1
				}
			case 3:
				{
					score += 2
				}
			}
		}
	case 2:
		{
			score += 3 + opt
		}
	case 3:
		{
			score += 6
			switch opt {
			case 1:
				{
					score += 2
				}
			case 2:
				{
					score += 3
				}
			case 3:
				{
					score += 1
				}
			}
		}
	}
	return score

}

func main() {
	data, err := ReadInput("inputs/input2.test")
	//data, err := ReadInput("input1.txt")
	helpers.HandleError(err)
	log.Println(data)
	// Rock    -> A  X
	// Paper   -> B  Y
	// Scissor -> C  Z
	base_points := []int{}
	total_points := 0
	// PART 1
	for _, j := range data {
		round := strings.Split(j, " ")
		points := AssignScore(round)
		base_points = append(base_points, points)
		total_points += points
	}
	log.Println(total_points)
	// Part 1: 10994

	// PART 2
	total_points = 0
	for _, j := range data {
		round := strings.Split(j, " ")
		total_points += GetMoveScore(round)
	}
	log.Println(total_points)
	// Part 2: 12526
}

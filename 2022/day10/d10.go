package main

import (
	"aoc2022/helpers"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	file, err := helpers.ReadInput("inputs/input10.txt")
	helpers.HandleError(err)
	x, cycle := 1, 0
	x_cycles := map[int]int{}
	signal_strength := 0
	for _, instruction := range file {
		move := strings.Split(instruction, " ")
		motion := move[0]
		count := 0
		if motion == "noop" {
			count = 0
			cycle++
			x_cycles[cycle] = x
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				signal_strength += cycle * x
			}
		} else {
			count, _ = strconv.Atoi(move[1])
			for i := 0; i < 2; i++ {
				cycle += 1
				x_cycles[cycle] = x
				if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
					signal_strength += cycle * x
				}
			}
			x += count
		}
	}
	log.Println(signal_strength)

	tube := []string{}
	row := ""
	pixel := 0
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			pixel++
			if math.Abs(float64(x_cycles[pixel])-float64(j)) <= 1 {
				row += "#"
			} else {
				row += "."
			}
		}
		tube = append(tube, "\n"+row)
		row = ""
	}
	log.Println(tube)
}

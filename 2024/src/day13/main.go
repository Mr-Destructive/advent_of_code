package main

import (
	"aoc2024"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func ParseMachines(machinesStr [][]string) [][]point {
	machines := [][]point{}
	for _, machineStr := range machinesStr {
		equations := []point{}
		for i, line := range machineStr {
			components := strings.Split(line, ", ")
			if len(components) < 2 {
				continue
			}
			xStr, yStr := "", ""
			if i == 0 {
				xStr = strings.TrimLeft(components[0], "Button A: X")
				yStr = strings.TrimLeft(components[1], "Y")
			} else if i == 1 {
				xStr = strings.TrimLeft(components[0], "Button B: X")
				yStr = strings.TrimLeft(components[1], "Y")
			} else {
				xStr = strings.TrimLeft(components[0], "Prize: X=")
				yStr = strings.TrimLeft(components[1], "Y=")
			}
			x, _ := strconv.Atoi(xStr)
			y, _ := strconv.Atoi(yStr)
			equations = append(equations, point{x, y})
		}
		machines = append(machines, equations)
	}
	return machines
}

func FindTokens(equations [][]point, offset int) int {
	score := 0
	for _, equation := range equations {
		a := equation[0]
		b := equation[1]
		c := equation[2]
		c.x += offset
		c.y += offset

		d := (a.x * b.y) - (a.y * b.x)
		x := ((b.x * (-c.y)) - (b.y * (-c.x)))
		y := ((a.y * (-c.x)) - (a.x * (-c.y)))
		if x%d == 0 && y%d == 0 {
			x = x / d
			y = y / d
			score += (3 * x) + y
		}
	}
	return score
}

func main() {
	lines := helpers.ReadFileLines("../inputs/day13/input.txt")
	machines := ParseMachines(lines)
	score1 := FindTokens(machines, 0)
	fmt.Println(score1)
	score2 := FindTokens(machines, 10000000000000)
	fmt.Println(score2)

}

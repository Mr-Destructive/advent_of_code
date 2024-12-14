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

func ParseRobots(lines []string) [][]point {
	robots := [][]point{}
	for _, line := range lines {
		robotStr := strings.Split(line, " ")
		robot := []point{}
		for _, rob := range robotStr {
			parts := strings.Split(strings.Trim(rob, "p=v="), ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			robot = append(robot, point{x, y})
		}
		robots = append(robots, robot)
	}
	return robots
}

func SimulateRobots(robots [][]point, width, height int) []point {
	movedRobots := []point{}
	for _, robot := range robots {
		coordinates := robot[0]
		velocity := robot[1]
		x := ((velocity.x * 100) + coordinates.x) % width
		y := ((velocity.y * 100) + coordinates.y) % height
		if x < 0 {
			x += width
		}
		if y < 0 {
			y += height
		}
		movedRobots = append(movedRobots, point{x, y})
	}
	return movedRobots
}

func PlaceQuadrants(robots []point, width, height int) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		quadrantX := width / 2
		quadrantY := height / 2
		switch {
		case robot.x == quadrantX || robot.y == quadrantY:
			continue
		case robot.x < quadrantX && robot.y < quadrantY:
			q1++
		case robot.x < quadrantX && robot.y > quadrantY:
			q2++
		case robot.y < height/2:
			q3++
		default:
			q4++
		}
	}
	return q1 * q2 * q3 * q4

}

func main() {
	lines := helpers.ReadFileLines("inputs/day14/prod.txt")
	robots := ParseRobots(lines)
	width, height := 101, 103
	movedPoints := SimulateRobots(robots, width, height)
	fmt.Println(movedPoints)
	score := PlaceQuadrants(movedPoints, width, height)
	fmt.Println(score)
}

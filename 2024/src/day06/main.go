package main

import (
	"bytes"
	"fmt"
	"os"
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

func ConstructGrid(lines []string) [][]string {
	grid := [][]string{}
	for _, line := range lines {
		row := []string{}
		for _, char := range strings.Split(line, "") {
			row = append(row, char)
		}
		grid = append(grid, row)
	}
	return grid
}

type point struct {
	x int
	y int
}

var directions []point = []point{
	point{0, -1}, // up
	point{0, 1},  // down
	point{1, 0},  // right
	point{-1, 0}, // left
}

func TraverseGrid(grid [][]string) (int, map[point]int) {
	score := 0
	startPoint := point{}
	for y, row := range grid {
		for x, char := range row {
			if char == "^" {
				startPoint = point{x, y}
			}
		}
	}
	visited := make(map[point]int)
	direction := directions[0]
	for {
		if _, ok := visited[startPoint]; !ok {
			score++
			visited[startPoint] = 1
		}
		isSafe, nextPoint, newDirection := moveNextPoint(startPoint, direction, grid)
		if !isSafe {
			fmt.Println(len(visited))
			return score, visited
		}
		startPoint = nextPoint
		direction = newDirection
	}
	return score, visited
}

func moveNextPoint(currPoint, direction point, grid [][]string) (bool, point, point) {
	isSafe, nextPoint, direction := move(currPoint, direction, grid)
	if !isSafe {
		return false, nextPoint, direction
	}
	if grid[nextPoint.y][nextPoint.x] == "#" {
		// turn 90
		if direction.y == 0 && direction.x == -1 {
			direction = directions[0]
		} else if direction.y == 0 && direction.x == 1 {
			direction = directions[1]
		} else if direction.y == -1 && direction.x == 0 {
			direction = directions[2]
		} else if direction.y == 1 && direction.x == 0 {
			direction = directions[3]
		}
		return moveNextPoint(currPoint, direction, grid)
	} else if grid[nextPoint.y][nextPoint.x] == "." || grid[nextPoint.y][nextPoint.x] == "^" {
		return true, nextPoint, direction
	}
	return false, currPoint, direction
}

func move(coordinate, direction point, grid [][]string) (bool, point, point) {
	nextPoint := point{coordinate.x + direction.x, coordinate.y + direction.y}
	if coordinate.x < 0 || coordinate.x >= len(grid) || coordinate.y < 0 || coordinate.y >= len(grid[0]) {
		return false, coordinate, direction
	}
	if nextPoint.y < 0 || nextPoint.y >= len(grid) || nextPoint.x < 0 || nextPoint.x >= len(grid[0]) {
		return false, coordinate, direction
	}
	return true, nextPoint, direction
}

func main() {
	lines := ReadFileLines("../../inputs/day06/input.txt")
	grid := ConstructGrid(lines)
	path, score := TraverseGrid(grid)
	fmt.Println(path)
	fmt.Println(score)
}

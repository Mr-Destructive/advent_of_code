package main

import (
	"bytes"
	"fmt"
	//"math"
	"os"
	"strings"
	"unicode"
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
		chars := []string{}
		for _, char := range strings.Split(line, "") {
			chars = append(chars, char)
		}
		grid = append(grid, chars)
	}
	return grid
}

func GetFrequencies(grid [][]string) map[string][][]int {
	frequencies := make(map[string][][]int)
	for x, line := range grid {
		for y, char := range line {
			if unicode.IsLetter(rune(char[0])) || unicode.IsDigit(rune(char[0])) {
				if _, ok := frequencies[char]; ok {
					frequencies[char] = append(frequencies[char], []int{x, y})
				} else {
					frequencies[char] = [][]int{{x, y}}
				}
			}
		}
	}
	return frequencies
}

func PlaceAnitnodes(grid [][]string, frequencies map[string][][]int) int {
	placedAntinodes := [][]int{}
	score := 0
	nonOverlappingnodes := make(map[string]bool)
	for frequency, posList := range frequencies {
		if len(posList) < 2 {
			continue
		}
		height, width := len(grid), len(grid[0])
		nNodes := len(posList)
		for i := 0; i < nNodes-1; i++ {
			for j := i + 1; j < nNodes; j++ {
				xdiff := posList[j][0] - posList[i][0]
				ydiff := posList[j][1] - posList[i][1]
				fmt.Println(posList[i][0], posList[i][1], frequency)
				fmt.Println(xdiff, ydiff, frequency)

				//if math.Abs(float64(xdiff)) >= 1 || math.Abs(float64(ydiff)) >= 1 {
				a1x := posList[i][0] - xdiff
				a1y := posList[i][1] - ydiff
				a2x := posList[j][0] + xdiff
				a2y := posList[j][1] + ydiff
				fmt.Printf("FOR %s -> a1: %d, %d == a2: %d, %d\n", frequency, a1x, a1y, a2x, a2y)
				if (a1x < width && a1x >= 0) && (a1y < height && a1y >= 0) {
					fmt.Println("A1 within bounds", a1x, a1y, posList[i], posList[j])
					placedAntinodes = append(placedAntinodes, []int{a1x, a1y})
				}
				if (a2x < width && a2x >= 0) && (a2y < height && a2y >= 0) {
					placedAntinodes = append(placedAntinodes, []int{a2x, a2y})
				}

			}
			for _, antinodes := range placedAntinodes {
				key := fmt.Sprintf("%d,%d", antinodes[0], antinodes[1])
				fmt.Println(key)
				if !nonOverlappingnodes[key] {
					nonOverlappingnodes[key] = true
					score++
				}
			}
			//}
		}
	}
	return score
}

func GetAntinodesWithinGrid(posList [][]int, i, j, width, height int) [][]int {
	placedAntinodes := [][]int{}
	x1, y1 := posList[i][0], posList[i][1]
	x2, y2 := posList[j][0], posList[j][1]

	if x1 < width && x1 >= 0 && y1 < height && y1 >= 0 {
		placedAntinodes = append(placedAntinodes, []int{x1, y1})

	}
	if x2 < width && x2 >= 0 && y2 < height && y2 >= 0 {
		placedAntinodes = append(placedAntinodes, []int{x2, y2})
	}

	return placedAntinodes
}

func IsWithinGrid(antenna []int, width, height int) bool {
	if antenna[0] < width && antenna[0] >= 0 && antenna[1] < height && antenna[1] >= 0 {
		return true
	}
	return false
}

func PlaceAnitnodesInline(grid [][]string, frequencies map[string][][]int) int {
	placedAntinodes := [][]int{}
	score := 0
	nonOverlappingnodes := make(map[string]bool)
	for _, posList := range frequencies {
		if len(posList) < 2 {
			continue
		}
		height, width := len(grid), len(grid[0])
		nNodes := len(posList)
		for i := 0; i < nNodes-1; i++ {
			for j := i + 1; j < nNodes; j++ {
				xdiff := posList[j][0] - posList[i][0]
				ydiff := posList[j][1] - posList[i][1]
				a1 := posList[i]
				a2 := posList[j]
				for IsWithinGrid(a1, width, height) || IsWithinGrid(a2, width, height) {
					if IsWithinGrid(a1, width, height) {
						placedAntinodes = append(placedAntinodes, a1)
					}
					if IsWithinGrid(a2, width, height) {
						placedAntinodes = append(placedAntinodes, a2)
					}
					a1 = []int{a1[0] - xdiff, a1[1] - ydiff}
					a2 = []int{a2[0] + xdiff, a2[1] + ydiff}

				}
				for _, antinodes := range placedAntinodes {
					key := fmt.Sprintf("%d,%d", antinodes[0], antinodes[1])
					if !nonOverlappingnodes[key] {
						nonOverlappingnodes[key] = true
						fmt.Println(key)
						score++
					}
				}
			}
		}
	}
	return score
}

func main() {
	lines := ReadFileLines("../../inputs/day08/prod.txt")
	grids := ConstructGrid(lines)
	frequencies := GetFrequencies(grids)
	//antinodes := PlaceAnitnodes(grids, frequencies)
	antinodes := PlaceAnitnodesInline(grids, frequencies)
	fmt.Println(antinodes)
}

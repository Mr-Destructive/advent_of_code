package main

import (
	"aoc2024"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func BlinkSlow(stones []string, blinks int) []string {
	for i := 0; i < blinks; i++ {
		index := 0
		for _, stone := range stones {
			stone = stones[index]
			if stone == "0" {
				stones[index] = "1"
			} else if len(stone)%2 == 0 {
				n := len(stone)
				first := stone[:n/2]
				second := stone[n/2:]
				x, _ := strconv.Atoi(first)
				y, _ := strconv.Atoi(second)
				stones[index] = strconv.Itoa(x)
				stones = slices.Insert(stones, index+1, strconv.Itoa(y))
				index += 1
			} else {
				n, _ := strconv.Atoi(stone)
				nStr := strconv.Itoa(n * 2024)
				stones[index] = nStr
			}
			index++
		}
	}
	return stones
}

func blinker(stone string, blinked map[string]int, blink int) int {
	if blink == 0 {
		return 1
	}
	blink--
	if s, ok := blinked[fmt.Sprintf("%s,%s", stone, blink)]; ok {
		return s
	}
	if stone == "0" {
		s := blinker("1", blinked, blink)
		blinked[fmt.Sprintf("%s,%s", stone, blink)] = s
		return s
	} else if len(stone)%2 == 0 {
		n := len(stone)
		first := stone[:n/2]
		second := stone[n/2:]
		x, _ := strconv.Atoi(first)
		y, _ := strconv.Atoi(second)
		s := blinker(strconv.Itoa(x), blinked, blink) + blinker(strconv.Itoa(y), blinked, blink)
		blinked[fmt.Sprintf("%s,%s", stone, blink)] = s
		return s
	}
	s, _ := strconv.Atoi(stone)
	s *= 2024
	count := blinker(strconv.Itoa(s), blinked, blink)
	blinked[fmt.Sprintf("%s,%s", stone, blink)] = count
	return count
}

func Blink(stones []string, blinks int) int {
	score := 0
	stoneBlinked := make(map[string]int)
	blinked := blinks
	for _, stone := range stones {
		score += blinker(stone, stoneBlinked, blinked)
	}
	return score
}

func CleanLines(lines []string) [][]string {
	stones := [][]string{}
	for _, line := range lines {
		stones = append(stones, strings.Split(line, " "))
	}
	return stones
}

func main() {
	lines := helpers.ReadFileLines("inputs/day11/prod.txt")
	stones := CleanLines(lines)
	count := Blink(stones[0], 75)
	fmt.Println(count)
}

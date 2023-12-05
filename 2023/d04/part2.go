package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	file_content := string(content)
	data := strings.Split(file_content, "\n")
	data = data[:len(data)-1]
	return data, err
}

func splitter(line string) int {
	count := 0
	parts := strings.Split(line, "|")
	parts[0] = strings.Trim(parts[0], " ")
	parts[1] = strings.Trim(parts[1], " ")
	ourCards := strings.Split(strings.Replace(strings.Split(parts[0], ":")[1], "  ", " ", -1), " ")
	winningCards := strings.Split(strings.Replace(parts[1], "  ", " ", -1), " ")

	for _, ourCard := range ourCards {
		for _, winner := range winningCards {
			if ourCard == winner {
				count += 1
			}
		}
	}
	return count
}

func totalScratchCards(winningCards, scratchCards []int) []int {
	for i := 0; i < len(winningCards); i++ {
		n := winningCards[i]
		for j := i + 1; j < i+1+n; j++ {
			scratchCards[j] += scratchCards[i]
		}
	}
	return scratchCards
}

func main() {
	lines, err := ReadInput("../inputs/input/d04.txt")
	if err != nil {
		panic(err)
	}
	winningCards := []int{}
	scratchCards := []int{}
	for i, line := range lines {
		c := splitter(line)
		fmt.Println(i+1, c)
		winningCards = append(winningCards, c)
		scratchCards = append(scratchCards, 1)
	}
	fmt.Println(winningCards)
	res := totalScratchCards(winningCards, scratchCards)
	count := 0
	for _, cards := range res {
		count += cards
	}
	fmt.Println(count)
}

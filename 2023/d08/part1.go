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

func splitter(lines []string) (string, map[string][]string) {
	nav := lines[0]
	maps := map[string][]string{}
	fmt.Println(nav)
	key := ""
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " = ")
		key = parts[0]
		val := strings.Split(strings.Trim(parts[1], "()"), ", ")
		right, left := val[0], val[1]
		maps[key] = []string{right, left}
	}
	return nav, maps
}

func runner(key string, maps map[string][]string) int {
	start := "AAA"
	e := "ZZZ"
	n := len(key)
	k := 0
	count := 1
	keyStr := string(key[k])
	for {
		l := maps[start]
		if keyStr == "R" {
			start = l[1]
		} else {
			start = l[0]
		}
		if start == e {
			break
		}
		k += 1
		if k >= n {
			k = 0
		}
		count += 1
		keyStr = string(key[k])
	}
	return count
}

func main() {
	lines, err := ReadInput("../inputs/input/d08.txt")
	if err != nil {
		panic(err)
	}
	key, maps := splitter(lines)
	fmt.Println(key, maps)
	count := runner(key, maps)
	fmt.Println(count)
}

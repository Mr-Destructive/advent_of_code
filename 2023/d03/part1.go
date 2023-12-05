package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadInput(file string) ([]string, error) {
	content, err := os.ReadFile(file)
	file_content := string(content)
	data := strings.Split(file_content, "\n")
	data = data[:len(data)-1]
	return data, err
}

func isSymbol(char rune) bool {
	return char != '.' && (char < '0' || char > '9')
}

func compareAdjacent(line string, lines []string, sum int) int {
	number := []int{}
	dig := ""
	n := len(line)
	isEngine := false

	for i, part := range line {
		if unicode.IsDigit(part) {
			dig += string(part)
		} else {
			dig = ""
		}

		if i+1 <= n-1 && unicode.IsDigit(part) && (isSymbol(rune(line[i+1])) || string(line[i+1]) == ".") {
			num, _ := strconv.Atoi(dig)
			number = append(number, num)
			dn := len(dig)
			x := len(lines)

			switch x {
			case 2:
				if line == lines[0] {
					// first line
					next := lines[1]
					s := i - len(dig) + 1
					if (s > 0 && isSymbol(rune(line[s-1]))) || (i+1 < n-1 && isSymbol(rune(line[i+1]))) {
						isEngine = true
						sum += num
						break
					}
					var nextStr string
					if s > 0 {
						nextStr = next[s-1 : s+dn+1]
					} else {
						nextStr = next[:s+dn+1]
					}
					for _, str := range nextStr {
						if isSymbol(rune(str)) {
							isEngine = true
							sum += num
							break
						}
					}
				} else {
					// last line
					prev := lines[0]
					s := i - len(dig) + 1
					if (s > 0 && isSymbol(rune(line[s-1]))) || (i+1 < n-1 && isSymbol(rune(line[i+1]))) {
						isEngine = true
						sum += num
						break
					}
					var prevStr string
					if s > 0 {
						prevStr = prev[s-1 : s+dn+1]
					} else {
						prevStr = prev[:s+dn+1]
					}
					for _, str := range prevStr {
						if isSymbol(rune(str)) {
							isEngine = true
							sum += num
							break
						}
					}
				}
			case 3:
				next, prev := lines[2], lines[0]
				s := i - len(dig) + 1
				if (s > 0 && isSymbol(rune(line[s-1]))) || (i+1 < n-1 && isSymbol(rune(line[i+1]))) {
					isEngine = true
					sum += num
					break
				}
				nextStr, prevStr := "", ""
				if s > 0 {
					nextStr = next[s-1 : s+dn+1]
					prevStr = prev[s-1 : s+dn+1]
				} else {
					nextStr = next[:s+dn+1]
					prevStr = prev[:s+dn+1]
				}
				for _, str := range prevStr {
					if isSymbol(rune(str)) {
						isEngine = true
						sum += num
						break
					}
				}
				for _, str := range nextStr {
					if isSymbol(rune(str)) {
						isEngine = true
						sum += num
						break
					}
				}
			}
		}
	}

	fmt.Println(number, isEngine)
	fmt.Println("SUM:", sum)
	return sum
}

func getEngineNumbers(lines []string) int {
	n := len(lines)
	sum := 0
	for i := 0; i < n; i++ {
		var adjs []string
		if i > 0 && i < n-1 && n > 2 {
			adjs = []string{lines[i-1], lines[i], lines[i+1]}
		} else if i == 0 && n > 1 {
			adjs = []string{lines[i], lines[i+1]}
		} else if i == n-1 && n > 1 {
			adjs = []string{lines[i-1], lines[i]}
		}
		sum = compareAdjacent(lines[i], adjs, sum)
	}
	return sum
}

func main() {
	lines, err := ReadInput("../inputs/input/d03.txt")
	if err != nil {
		panic(err)
	}
	sum := getEngineNumbers(lines)
	fmt.Println(sum)
}

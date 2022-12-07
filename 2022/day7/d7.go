package main

import (
	"aoc2022/helpers"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("inputs/input7.txt")
	helpers.HandleError(err)
	contents := []string{}
	contents = append(contents, strings.Split(string(file), "\n$ ")...)
	dirs_content := map[string]int{}
	for i := 0; i < len(contents); i++ {
		s := strings.Split(contents[i], "\n")
		if len(s) > 1 {
			size := 0
			for _, f := range s {
				if len(f) > 2 {
					output := strings.Split(f, " ")
					file_size, _ := strconv.Atoi(output[0])
					size += file_size
				}
			}
			if size <= 100000 {
				dirs_content[string(contents[i-1][3:])] += size
			}
		}
	}
	// here
	for i := 0; i < len(contents); i++ {
		s := strings.Split(contents[i], "\n")
		if len(s) > 1 {
			for _, f := range s {
				if len(f) > 2 {
					output := strings.Split(f, " ")
					if output[0] == "dir" {
						if string(contents[i-1][4:]) != " /" {
							dirs_content[string(contents[i-1][3:])] += dirs_content[output[1]]
						}
					}
				}
			}
		}
	}
	// here

	for k, v := range dirs_content {
		if v > 100000 {
			delete(dirs_content, k)
		}
	}
	log.Println(dirs_content)
	sum := 0
	for _, v := range dirs_content {
		sum += v
	}
	log.Println(sum)
}

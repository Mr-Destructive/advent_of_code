package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.ReadFile("input4.txt")
	HandleError(err)
	file_content := strings.Split(string(file), "\n")
	file_content = file_content[:len(file_content)-1]
	part_1, part_2 := 0, 0
	for _, assignments := range file_content {
		assgmnts := strings.Split(assignments, ",")
		first_pair := strings.Split(assgmnts[0], "-")
		second_pair := strings.Split(assgmnts[1], "-")
		// (x1, y1)    (x2, y2)
		x1, _ := strconv.ParseFloat(first_pair[0], 64)
		x2, _ := strconv.ParseFloat(second_pair[0], 64)
		y1, _ := strconv.ParseFloat(first_pair[1], 64)
		y2, _ := strconv.ParseFloat(second_pair[1], 64)
		if (x1 <= x2 && y1 >= y2) || (x2 <= x1 && y2 >= y1) {
			part_1 += 1
		}
		if x1 <= y2 && x2 <= y1 {
			part_2 += 1
		}
	}
	log.Println(part_1) // 509
	log.Println(part_2) // 870
}

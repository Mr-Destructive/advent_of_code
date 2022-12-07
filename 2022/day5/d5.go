package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"aoc2022/helpers"
)

func Get_Move_Commands(cmd string) (int, int, int) {
	chunks := strings.Split(cmd, " ")
	number_of_crates, _ := strconv.Atoi(chunks[1])
	from_crate, _ := strconv.Atoi(chunks[3])
	to_crate, _ := strconv.Atoi(chunks[5])
	return number_of_crates, from_crate - 1, to_crate - 1
}

func reverse(array []string) []string {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func main() {
	file, err := os.ReadFile("inputs/input5.test")
	file_content := strings.Split(string(file), "\n\n")
	helpers.HandleError(err)
	number_of_stacks := 3
	lines := strings.Split(file_content[0], "\n")
	//stacks := len(strings.Split(lines[len(lines)-3], "  "))
	stack_list := [][]string{}
	for i := 0; i < number_of_stacks; i += 1 {
		stack := []string{}
		for j := 1; j < len(lines[i])-1; j += 4 {
			stack = append(stack, string(lines[i][j]))
		}
		stack_list = append(stack_list, stack)
	}
	stack_list = stack_list[:number_of_stacks] //-1]
	log.Println(stack_list)
	new_stack_list := [][]string{}
	new_stack_list = append(new_stack_list, stack_list...)
	new_stack_list = transpose(new_stack_list)

	log.Println(new_stack_list)
	for i := 0; i < len(new_stack_list); i++ {
		for j := 0; j < len(new_stack_list[i]); j++ {
			//log.Println(new_stack_list[i][j], new_stack_list[j][i])
			if new_stack_list[i][j] == " " {
				new_stack_list[i] = append(new_stack_list[i][:j], new_stack_list[i][j+1:]...)
				i = 0
			}
		}
	}
	log.Println(new_stack_list)

	commands := strings.Split(file_content[1], "\n")
	commands = commands[:len(commands)-1]
	for _, line := range commands {
		crates, from, to := Get_Move_Commands(string(line))
		//for n := 0; n < crates; n++ {
		//last_elem := len(new_stack_list[from]) - 1
		reversed_crates := reverse(new_stack_list[from][:crates])
		new_stack_list[to] = append(new_stack_list[to], new_stack_list[from][:crates]...)
		//new_stack_list[to] = append(reversed_crates, new_stack_list[to]...)
		if len(new_stack_list[from]) > 0 {
			log.Println(reversed_crates)
			//new_stack_list[from] = append(new_stack_list[from][:crates-1], new_stack_list[from][crates:]...)
			new_stack_list[from] = append(new_stack_list[from][:crates-1], new_stack_list[from][crates:]...)
		} else {
			//new_stack_list[from] = new_stack_list[from][:crates-1]
			new_stack_list[from] = append(new_stack_list[from], reversed_crates...)
		}
		//}
		log.Println(crates, new_stack_list)
	}
	log.Println(new_stack_list)
	for _, line := range new_stack_list {
		end := len(line) - 1
		print(line[end])
	}
}

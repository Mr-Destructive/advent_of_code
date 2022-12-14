package main

import (
	"aoc2022/helpers"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parse_lines(lines []string) (int, []string, string, string, string, int, int, int) {
	var monkey_no int
	var divisible_by int
	fmt.Sscanf(lines[0], "Monkey %d:", &monkey_no)
	worries := strings.Split(lines[1], "  Starting items: ")
	r, l, operator := "", "", ""
	true_monkey, false_monkey := 0, 0
	fmt.Sscanf(lines[2], "  Operation: new = %s %s %s", &l, &operator, &r)
	fmt.Sscanf(lines[3], "  Test: divisible by %d", &divisible_by)
	fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &true_monkey)
	fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &false_monkey)
	return monkey_no, worries, l, operator, r, divisible_by, true_monkey, false_monkey
}

func parse_new_worry(worry int, l, operator, r string) int {
	var l_worry int
	var r_worry int
	var new_worry int
	if l == "old" {
		l_worry = worry
	} else {
		num, _ := strconv.Atoi(l)
		l_worry = num
	}
	if r == "old" {
		r_worry = worry
	} else {
		num, _ := strconv.Atoi(r)
		r_worry = num
	}
	if operator == "+" {
		new_worry = l_worry + r_worry
	} else if operator == "-" {
		new_worry = l_worry - r_worry
	} else if operator == "*" {
		new_worry = l_worry * r_worry
	} else {
		new_worry = l_worry / r_worry
	}
	return new_worry
}

func monkey_bussiness(lines []string, monkey_items [][]int) ([][]int, int) {
	var i int = 0
	var worry_level []int
	var monkey_no int
	var inspected_items int = 0
	fmt.Sscanf(lines[0], "Monkey %d:", &monkey_no)
	monkey_no, _, l, operator, r, divisible_by, true_monkey, false_monkey := parse_lines(lines)
	for len(monkey_items[monkey_no]) > 0 {
		inspected_items += 1
		i = len(monkey_items[monkey_no]) - 1
		worry := monkey_items[monkey_no][i]
		i -= 1
		worry_level = append(worry_level, worry)
		new_worry := parse_new_worry(worry, l, operator, r)
		new_worry /= 3
		for i, elem := range monkey_items[monkey_no] {
			if elem == worry {
				if len(monkey_items[monkey_no]) == 1 {
					monkey_items[monkey_no] = monkey_items[monkey_no][:len(monkey_items[monkey_no])]
				}
				monkey_items[monkey_no] = append(monkey_items[monkey_no][:i], monkey_items[monkey_no][i+1:]...)
				break
			}
			last := len(monkey_items[monkey_no]) - 1
			worry = monkey_items[monkey_no][last]
		}
		if new_worry%divisible_by == 0 {
			monkey_items[true_monkey] = append(monkey_items[true_monkey], new_worry)
		} else {
			monkey_items[false_monkey] = append(monkey_items[false_monkey], new_worry)
		}
	}
	return monkey_items, inspected_items
}
func monkey_bussiness_2(lines []string, monkey_items [][]int) ([][]int, int) {
	var i int = 0
	//var worry_level []int
	var monkey_no int
	var inspected_items int = 0
	fmt.Sscanf(lines[0], "Monkey %d:", &monkey_no)
	monkey_no, _, l, operator, r, divisible_by, true_monkey, false_monkey := parse_lines(lines)
	for len(monkey_items[monkey_no]) > 0 {
		inspected_items += 1
		i = len(monkey_items[monkey_no]) - 1
		worry := monkey_items[monkey_no][i]
		i -= 1
		//worry_level = append(worry_level, worry)
		new_worry := parse_new_worry(worry, l, operator, r)
		for i, elem := range monkey_items[monkey_no] {
			if elem == worry {
				if len(monkey_items[monkey_no]) == 1 {
					monkey_items[monkey_no] = monkey_items[monkey_no][:len(monkey_items[monkey_no])]
				}
				monkey_items[monkey_no] = append(monkey_items[monkey_no][:i], monkey_items[monkey_no][i+1:]...)
				break
			}
			last := len(monkey_items[monkey_no]) - 1
			worry = monkey_items[monkey_no][last]
		}
		if new_worry%divisible_by == 0 {
			monkey_items[true_monkey] = append(monkey_items[true_monkey], new_worry)
		} else {
			monkey_items[false_monkey] = append(monkey_items[false_monkey], new_worry)
		}
	}
	return monkey_items, inspected_items
}

func main() {
	file, err := helpers.ReadInputTrim("inputs/input11.test", "\n\n")
	helpers.HandleError(err)
	var monkey_items = [][]int{}
	var monkey_items_2 = [][]int{}
	inspected_items := []int{}
	inspected_items_2 := []int{}
	for i := 0; i < len(file); i++ {
		monkey_items = append(monkey_items, []int{})
		monkey_items_2 = append(monkey_items_2, []int{})
		inspected_items = append(inspected_items, 0)
		inspected_items_2 = append(inspected_items_2, 0)
	}
	for m := 0; m < 1000; m++ {
		inspected_item := 0
		inspected_item_2 := 0
		for k, s := range file {
			lines := strings.Split(s, "\n")
			var monkey_no int
			if m == 0 {
				worries := strings.Split(lines[1], "  Starting items: ")
				worry := strings.Split(worries[1], ", ")
				worry_list := []int{}
				for _, w := range worry {
					temp, _ := strconv.Atoi(w)
					worry_list = append(worry_list, temp)

				}
				fmt.Sscanf(lines[0], "Monkey %d:", &monkey_no)
				for _, w := range worry_list {
					monkey_items[monkey_no] = append(monkey_items[monkey_no], w)
					monkey_items_2[monkey_no] = append(monkey_items_2[monkey_no], w)
				}
			}
			monkey_items, inspected_item = monkey_bussiness(lines, monkey_items)
			monkey_items_2, inspected_item_2 = monkey_bussiness_2(lines, monkey_items_2)
			inspected_items[k] += inspected_item
			inspected_items_2[k] += inspected_item_2
			if m == 19 || m == 999 {
				fmt.Println(m, "round ->", k, " -> ", inspected_item_2, inspected_items_2[k])
			}
		}
		for i := 0; i < 4; i++ {
			//fmt.Println(m, i, inspected_items_2[i])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspected_items_2)))
	log.Println(inspected_items_2[0] * inspected_items_2[1])
}

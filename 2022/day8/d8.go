package main

import (
	"aoc2022/helpers"
	"log"
)

func check_trees(file []string, x, y, rows int) bool {
	ur, dr, rc, lc := 0, 0, 0, 0
	for i := 0; i < rows; i++ {
		if i > x {
			if file[x][y] > file[i][y] {
				dr++
			}
		} else if x == i {
			if ur == x {
				return true
			} else {
				ur = 0
			}
		} else {
			if file[x][y] > file[i][y] {
				ur++
			}
		}
	}
	for j := 0; j < rows; j++ {
		if j > y {
			if file[x][y] > file[x][j] {
				rc++
			}
		} else if y == j {
			if lc == y {
				return true
			}
			lc = 0
		} else {
			if file[x][y] > file[x][j] {
				lc++
			}
		}
	}
	if dr+1 == rows-x || rc+1 == rows-y {
		return true
	} else {
		return false
	}
}
func scenic_score(file []string, x, y, rows int) int {
	ur, dr, rc, lc := 0, 0, 0, 0
	scenic_score := 0
	block_r, block_l, block_u, block_d := false, false, false, false
	for i := x - 1; i >= 0; i-- {
		if file[x][y] > file[i][y] && !block_u {
			ur++
		} else if file[x][y] <= file[i][y] && !block_u {
			ur++
			block_u = true
		}
	}
	for i := x + 1; i < rows; i++ {
		if file[x][y] > file[i][y] && !block_d {
			dr++
		} else if file[x][y] <= file[i][y] && !block_d {
			dr++
			block_d = true
		}
	}
	for j := y - 1; j >= 0; j-- {
		if file[x][y] > file[x][j] && !block_l {
			lc++
		} else if file[x][y] <= file[x][j] && !block_l {
			lc++
			block_l = true
		}
	}
	for j := y + 1; j < rows; j++ {
		if file[x][y] > file[x][j] && !block_r {
			rc++
		} else if file[x][y] <= file[x][j] && !block_r {
			rc++
			block_r = true
		}
	}
	scenic_score = ur * dr * rc * lc
	return scenic_score
}

func main() {
	file, err := helpers.ReadInput("inputs/input8.txt")
	helpers.HandleError(err)
	rows := len(file)
	total_trees := 0
	inside_trees := 0
	border_count := (rows * rows) - ((rows - 2) * (rows - 2))
	scenic_score_list := []int{}
	max_scenic_score := 0
	for n := 1; n < rows-1; n++ {
		for i := 1; i < rows-1; i++ {
			//log.Println("f", n, i, string(file[n][i]))
			val := check_trees(file, n, i, rows)
			if val {
				score := scenic_score(file, n, i, rows)
				scenic_score_list = append(scenic_score_list, score)
				if max_scenic_score < score {
					max_scenic_score = score
				}
				inside_trees++
			}
		}
	}
	total_trees = inside_trees + border_count
	log.Println(total_trees)
	log.Println(scenic_score_list)
	log.Println(max_scenic_score)
}

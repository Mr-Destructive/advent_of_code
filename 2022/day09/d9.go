package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"aoc2022/helpers"
)

type point struct {
	x int
	y int
}

func Move(move string, p point) point {
	switch move {
	case "R":
		p.x += 1
	case "L":
		p.x -= 1
	case "U":
		p.y += 1
	case "D":
		p.y -= 1
	}
	return p
}

func (p point) MapMove(move string) point {
	p = Move(move, p)
	return p
}

func (p point) FollowHead(head point) point {
	distance := math.Max(math.Abs(float64(p.x-head.x)), math.Abs(float64(p.y-head.y)))
	if distance > 1 {
		tail_x := head.x - p.x
		tail_y := head.y - p.y
		if math.Abs(float64(tail_x)) == 2 {
			tail_x /= 2
		}
		if math.Abs(float64(tail_y)) == 2 {
			tail_y /= 2
		}
		p.x += tail_x
		p.y += tail_y
	}
	return p
}

func fillMatrix(p map[point]bool, x, y int) map[point]bool {
	if _, ok := p[point{x, y}]; !ok {
		p[point{x, y}] = true
	}
	return p
}

func main() {
	file, err := helpers.ReadInput("inputs/input9.txt")
	helpers.HandleError(err)
	head, tail := point{0, 0}, point{0, 0}
	matrix := map[point]bool{{0, 0}: true}
	for _, m := range file {
		move := strings.Split(m, " ")
		motion := move[0]
		distance, _ := strconv.Atoi(move[1])
		for i := 0; i < distance; i++ {
			head = head.MapMove(motion)
			tail = tail.FollowHead(head)
			matrix = fillMatrix(matrix, tail.x, tail.y)
		}

	}
	log.Println(head, tail)
	log.Println(len(matrix))

	knots := []point{}
	matrix = map[point]bool{{0, 0}: true}
	for i := 0; i < 10; i++ {
		knots = append(knots, point{0, 0})
	}
	for _, m := range file {
		move := strings.Split(m, " ")
		motion := move[0]
		distance, _ := strconv.Atoi(move[1])
		for i := 0; i < distance; i++ {
			knots[0] = knots[0].MapMove(motion)
			for j := 1; j < len(knots); j++ {
				knots[j] = knots[j].FollowHead(knots[j-1])
			}
			tail = knots[len(knots)-1]
			matrix = fillMatrix(matrix, tail.x, tail.y)
		}
	}
	log.Println(len(matrix))
}

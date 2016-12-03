package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const input = "R2, L5, L4, L5, R4, R1, L4, R5, R3, R1, L1, L1, R4, L4, L1, R4, L4, R4, L3, R5, R4, R1, R3, L1, L1, R1, L2, R5, L4, L3, R1, L2, L2, R192, L3, R5, R48, R5, L2, R76, R4, R2, R1, L1, L5, L1, R185, L5, L1, R5, L4, R1, R3, L4, L3, R1, L5, R4, L4, R4, R5, L3, L1, L2, L4, L3, L4, R2, R2, L3, L5, R2, R5, L1, R1, L3, L5, L3, R4, L4, R3, L1, R5, L3, R2, R4, R2, L1, R3, L1, L3, L5, R4, R5, R2, R2, L5, L3, L1, L1, L5, L2, L3, R3, R3, L3, L4, L5, R2, L1, R1, R3, R4, L2, R1, L1, R3, R3, L4, L2, R5, R5, L1, R4, L5, L5, R1, L5, R4, R2, L1, L4, R1, L1, L1, L5, R3, R4, L2, R1, R2, R1, R1, R3, L5, R1, R4"

//const input = "R8, R4, R4, R8"

const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)

type player struct {
	pos     coord
	dir     int
	visited map[coord]bool
}

type coord struct {
	x, y int
}

func (c *coord) distance() int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func main() {
	inputList := strings.Split(input, ", ")
	p := player{}
	p.visited = make(map[coord]bool)
	for _, v := range inputList {
		if p.step(v) {
			fmt.Println("Solution is:", p.pos.distance())
			return
		}
	}
}

func (p *player) step(input string) bool {
	p.rotate(input)

	length, _ := strconv.Atoi(input[1:])
	for i := 0; i < length; i++ {
		if p.visited[p.pos] {
			return true
		}
		p.visited[p.pos] = true
		p.move(1)
	}
	return false
}

func (p *player) move(length int) {
	switch p.dir {
	case NORTH:
		p.pos.y += length
		return
	case EAST:
		p.pos.x += length
		return
	case SOUTH:
		p.pos.y -= length
		return
	case WEST:
		p.pos.x -= length
		return
	}
}

func (p *player) rotate(input string) {
	dirString := input[:1]
	if dirString == "R" {
		p.dir = (p.dir + 1) % 4
	} else {
		p.dir = p.dir - 1
		if p.dir < 0 {
			p.dir = WEST
		}
	}
}

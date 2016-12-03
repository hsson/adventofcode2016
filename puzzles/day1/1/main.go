package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const input = "R2, L5, L4, L5, R4, R1, L4, R5, R3, R1, L1, L1, R4, L4, L1, R4, L4, R4, L3, R5, R4, R1, R3, L1, L1, R1, L2, R5, L4, L3, R1, L2, L2, R192, L3, R5, R48, R5, L2, R76, R4, R2, R1, L1, L5, L1, R185, L5, L1, R5, L4, R1, R3, L4, L3, R1, L5, R4, L4, R4, R5, L3, L1, L2, L4, L3, L4, R2, R2, L3, L5, R2, R5, L1, R1, L3, L5, L3, R4, L4, R3, L1, R5, L3, R2, R4, R2, L1, R3, L1, L3, L5, R4, R5, R2, R2, L5, L3, L1, L1, L5, L2, L3, R3, R3, L3, L4, L5, R2, L1, R1, R3, R4, L2, R1, L1, R3, R3, L4, L2, R5, R5, L1, R4, L5, L5, R1, L5, R4, R2, L1, L4, R1, L1, L1, L5, R3, R4, L2, R1, R2, R1, R1, R3, L5, R1, R4"

const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)

type player struct {
	x, y, dir int
}

func main() {
	inputList := strings.Split(input, ", ")
	p := player{}
	for _, v := range inputList {
		p.rotate(v)
		p.move(v)
	}
	fmt.Println("Solution is:", math.Abs(float64(p.x))+math.Abs(float64(p.y)))
}

func (p *player) move(input string) {
	length, _ := strconv.Atoi(input[1:])
	switch p.dir {
	case NORTH:
		p.y += length
		return
	case EAST:
		p.x += length
		return
	case SOUTH:
		p.y -= length
		return
	case WEST:
		p.x -= length
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

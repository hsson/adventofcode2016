package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

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
	input, _ := ioutil.ReadFile("../input.txt")
	inputList := strings.Split(string(input), ", ")
	p := player{}
	for _, v := range inputList {
		p.rotate(v)
		p.move(v)
	}
	fmt.Println(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
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

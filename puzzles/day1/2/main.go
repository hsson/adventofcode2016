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
	pos     coord
	dir     int
	visited map[coord]bool
}

type coord struct {
	x, y int
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	inputList := strings.Split(string(input), ", ")
	p := player{}
	p.visited = make(map[coord]bool)
	for _, v := range inputList {
		if p.step(v) {
			fmt.Println(int(math.Abs(float64(p.pos.x)) + math.Abs(float64(p.pos.y))))
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

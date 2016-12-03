package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type player struct {
	pos coord
}

type coord struct {
	x, y int
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	p := player{}
	p.pos = coord{1, 1} // Start on number 5
	for _, l := range lines {
		fmt.Print(p.process(l))
	}
}

func (p *player) process(line string) string {
	for _, c := range line {
		p.pos.move(c)
	}
	return strconv.Itoa(p.pos.y*3 + (p.pos.x + 1))
}

func (c *coord) move(dir rune) {
	if dir == 'U' {
		c.y--
		if c.y < 0 {
			c.y = 0
		}
	} else if dir == 'D' {
		c.y++
		if c.y == 3 {
			c.y = 3 - 1
		}
	} else if dir == 'L' {
		c.x--
		if c.x < 0 {
			c.x = 0
		}
	} else {
		c.x++
		if c.x == 3 {
			c.x = 3 - 1
		}
	}
}

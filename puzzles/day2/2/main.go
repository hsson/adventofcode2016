package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var keypad = [][]rune{
	[]rune{' ', ' ', '1', ' ', ' '},
	[]rune{' ', '2', '3', '4', ' '},
	[]rune{'5', '6', '7', '8', '9'},
	[]rune{' ', 'A', 'B', 'C', ' '},
	[]rune{' ', ' ', 'D', ' ', ' '},
}

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
	p.pos = coord{0, 2} // Start at number 5
	for _, line := range lines {
		fmt.Printf("%c", p.process(line))
	}
}

func (p *player) process(line string) rune {
	for _, c := range line {
		p.pos.move(c)
	}
	return keypad[p.pos.y][p.pos.x]
}

func (c *coord) move(dir rune) {
	if dir == 'U' && c.y > 0 {
		nY := c.y - 1
		if keypad[nY][c.x] != ' ' {
			c.y = nY
		}
	} else if dir == 'D' && c.y < 4 {
		nY := c.y + 1
		if keypad[nY][c.x] != ' ' {
			c.y = nY
		}
	} else if dir == 'L' && c.x > 0 {
		nX := c.x - 1
		if keypad[c.y][nX] != ' ' {
			c.x = nX
		}
	} else if dir == 'R' && c.x < 4 {
		nX := c.x + 1
		if keypad[c.y][nX] != ' ' {
			c.x = nX
		}
	}
}

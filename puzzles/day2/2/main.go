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

type coord struct {
	x, y int
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	pos := coord{0, 2} // Start at number 5
	for _, line := range lines {
		for _, c := range line {
			pos.move(c)
		}
		fmt.Printf("%c", keypad[pos.y][pos.x])
	}
}

func (c *coord) move(dir rune) {
	if dir == 'U' && c.y > 0 && keypad[c.y-1][c.x] != ' ' {
		c.y--
	} else if dir == 'D' && c.y < 4 && keypad[c.y+1][c.x] != ' ' {
		c.y++
	} else if dir == 'L' && c.x > 0 && keypad[c.y][c.x-1] != ' ' {
		c.x--
	} else if dir == 'R' && c.x < 4 && keypad[c.y][c.x+1] != ' ' {
		c.x++
	}
}

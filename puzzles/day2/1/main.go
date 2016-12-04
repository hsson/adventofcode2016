package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	pos := coord{1, 1} // Start on number 5
	for _, l := range lines {
		for _, c := range l {
			pos.move(c)
		}
		fmt.Print(strconv.Itoa(pos.y*3 + (pos.x + 1)))
	}
}

func (c *coord) move(dir rune) {
	if dir == 'U' && c.y > 0 {
		c.y--
	} else if dir == 'D' && c.y < 2 {
		c.y++
	} else if dir == 'L' && c.x > 0 {
		c.x--
	} else if dir == 'R' && c.x < 2 {
		c.x++
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	count := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < len(lines); j += 3 {
			x := parseRow(lines[j], i)
			y := parseRow(lines[j+1], i)
			z := parseRow(lines[j+2], i)
			if x+y > z && y+z > x && x+z > y {
				count++
			}
		}
	}
	fmt.Println(count)
}

func parseRow(l string, col int) int {
	l = strings.TrimSpace(l)
	l = strings.Replace(l, "    ", " ", -1)
	l = strings.Replace(l, "   ", " ", -1)
	l = strings.Replace(l, "  ", " ", -1)
	ll := strings.Split(l, " ")
	x, _ := strconv.Atoi(ll[col])
	return x
}

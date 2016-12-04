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
	for _, l := range lines {
		x, y, z := parseTriangle(l)
		if x+y > z && y+z > x && x+z > y {
			count++
		}
	}
	fmt.Println(count)
}

func parseTriangle(l string) (x, y, z int) {
	l = strings.TrimSpace(l)
	l = strings.Replace(l, "    ", " ", -1)
	l = strings.Replace(l, "   ", " ", -1)
	l = strings.Replace(l, "  ", " ", -1)
	ll := strings.Split(l, " ")
	x, _ = strconv.Atoi(ll[0])
	y, _ = strconv.Atoi(ll[1])
	z, _ = strconv.Atoi(ll[2])
	return
}

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	line := string(input)
	output := ""
	for line != "" {
		char := string(line[0])
		if char != "(" {
			output += char
			line = line[1:]
		} else {
			var x int
			var y int
			x, y, line = processMarker(line)
			for i := 0; i < y; i++ {
				output += line[0:x]
			}
			line = line[x:]
		}
	}
	fmt.Println(len(output))
}

func processMarker(line string) (int, int, string) {
	count := 0
	for string(line[count]) != ")" {
		count++
	}
	XY := strings.Split(line[1:count], "x")
	x, _ := strconv.Atoi(XY[0])
	y, _ := strconv.Atoi(XY[1])

	return x, y, line[count+1:]
}

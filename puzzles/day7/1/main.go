package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var inBrackets = regexp.MustCompile(`\[(.*?)\]`)
var outsideBrackets = regexp.MustCompile(`(.*?)(?:\[.*?\]|$)`)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	supported := 0
	for _, line := range lines {
		if isLineValid(extractSubparts(line)) {
			supported++
		}
	}
	fmt.Println(supported)
}

func isLineValid(normals, hypernets []string) bool {
	for _, hypernet := range hypernets {
		if containsABBA(hypernet) {
			return false
		}
	}
	for _, normal := range normals {
		if containsABBA(normal) {
			return true
		}
	}
	return false
}

func containsABBA(line string) bool {
	for i := 0; i < len(line)-3; i++ {
		partOne := line[i : i+2]
		partTwo := line[i+2 : i+4]
		if partOne[1] == partTwo[0] && partOne[0] == partTwo[1] && partOne[0] != partOne[1] {
			return true
		}
	}
	return false
}

func extractSubparts(line string) (normals, hypernets []string) {
	hypernets = inBrackets.FindAllString(line, -1)
	for i := range hypernets {
		// Remove brackets
		hypernets[i] = hypernets[i][1 : len(hypernets[i])-1]
	}
	for _, val := range outsideBrackets.FindAllStringSubmatch(line, -1) {
		normals = append(normals, val[1])
	}
	return
}

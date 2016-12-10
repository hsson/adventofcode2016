package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var inBrackets = regexp.MustCompile(`\[(.*?)\]`)
var outsideBrackets = regexp.MustCompile(`(.*?)\[.*?\]`)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	supported := 0
	for _, line := range lines {
		break
		if isLineValid(extractSubparts(line)) {
			supported++
		}
		break
	}
	isLineValid(extractSubparts("hhttjuhgvcgkisaqof[obpleewrfrrsgpumz]umcmeaytqjlqkyrawp[rhkhciyzmxciiysv]"))
	fmt.Println(supported)
}

func isLineValid(normals, hypernets []string) bool {
	return true
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

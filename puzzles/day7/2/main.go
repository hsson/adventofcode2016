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
	var abas []string
	for _, norm := range normals {
		abas = append(abas, getABAs(norm)...)
	}
	for _, hyper := range hypernets {
		for _, aba := range abas {
			if containsBAB(hyper, aba) {
				return true
			}
		}
	}
	return false
}

func containsBAB(line, aba string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == aba[1] && line[i+2] == aba[1] && line[i+1] == aba[0] {
			return true
		}
	}
	return false
}

func getABAs(line string) []string {
	var abas []string
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] && line[i] != line[i+1] {
			abas = append(abas, line[i:i+3])
		}
	}
	return abas
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

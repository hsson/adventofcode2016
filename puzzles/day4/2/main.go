package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	aOffset   = 97
	alphaSize = 26
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		id, _ := strconv.Atoi(line[len(line)-10 : len(line)-7])
		enc := line[:len(line)-10]
		if e := unencrypt(enc, id); strings.Contains(e, "north") {
			fmt.Printf("%v | %s\n", id, e)
		}
	}
}

func unencrypt(enc string, id int) string {
	new := ""
	for _, l := range enc {
		if l == '-' {
			new += " "
		} else {
			new += string(((l + rune(id) - aOffset) % alphaSize) + aOffset)
		}
	}
	return new
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	occurances := make(map[int][]string, len(lines))
	message := make([]string, 8)
	for _, line := range lines {
		for i, c := range line {
			occurances[i] = append(occurances[i], string(c))
		}
	}
	for key, val := range occurances {
		message[key] = getMostOccuring(val)
	}
	fmt.Println(strings.Join(message, ""))
}

func getMostOccuring(list []string) string {
	m := make(map[string]int, len(list))
	max := -1
	maxChar := ""
	for _, char := range list {
		m[char]++
	}
	for key, val := range m {
		if val > max {
			max = val
			maxChar = key
		}
	}
	return maxChar
}

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
		message[key] = getLeastOccuring(val)
	}
	fmt.Println(strings.Join(message, ""))
}

func getLeastOccuring(list []string) string {
	m := make(map[string]int, len(list))
	least := 9999999
	leastChar := ""
	for _, char := range list {
		m[char]++
	}
	for key, val := range m {
		if val < least {
			least = val
			leastChar = key
		}
	}
	return leastChar
}

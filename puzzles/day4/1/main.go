package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("../input.txt")
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, line := range lines {
		check := line[len(line)-7:][1:6]
		id, _ := strconv.Atoi(line[len(line)-10 : len(line)-7])
		enc := strings.Replace(line[:len(line)-10], "-", "", -1)
		if isValid(enc, check) {
			sum += id
		}
	}
	fmt.Println(sum)
}

func isValid(enc, check string) bool {
	m := make(map[string]int)
	for _, c := range strings.Split(enc, "") {
		m[c] = m[c] + 1
	}
	max := -1
	for _, v := range m {
		if v >= max {
			max = v
		}
	}
	occ := make([]string, max)
	for k, v := range m {
		occ[v-1] += k
	}

	checked := 0
	top := ""
	for i := len(occ) - 1; i >= 0; i-- {
		if occ[i] == "" {
			continue
		}
		letters := strings.Split(occ[i], "")
		sort.Strings(letters)
		for _, l := range letters {
			if checked >= len(check) {
				return top == check
			}
			top += l
			checked++
		}
	}
	return top == check
}

package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const input = "uqwqemis"

func main() {
	var count int64
	count = 0
	passLength := 0
	password := make([]string, 8)
	for passLength < 8 {
		data := []byte(input + strconv.FormatInt(count, 10))
		hashed := fmt.Sprintf("%x", md5.Sum(data))
		if hashed[:5] == "00000" {
			pos, err := strconv.Atoi(string(hashed[5]))
			if err == nil && pos < 8 && password[pos] == "" {
				char := string(hashed[6])
				password[pos] = char
				passLength++
				fmt.Printf("c: %v, pos: %v, count:%v\n", char, pos, count)
			}
		}
		count++
	}
	for _, c := range password {
		fmt.Print(c)
	}
	fmt.Println()
}

package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const input = "uqwqemis"

func main() {
	count := 0
	password := ""
	for len(password) < 8 {
		data := []byte(input + strconv.Itoa(count))
		hashed := fmt.Sprintf("%x", md5.Sum(data))
		if hashed[:5] == "00000" {
			char := string(hashed[5])
			password += char
			fmt.Print(char)
		}
		count++
	}
	fmt.Println()
}

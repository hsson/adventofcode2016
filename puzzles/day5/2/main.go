package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

const input = "uqwqemis"

func main() {
	var count int64
	count = 0
	passLength := 0
	password := strings.Split("________", "")
	fmt.Printf("\rDECRYPTING: %v", strings.Join(password, ""))
	for passLength < 8 {
		data := []byte(input + strconv.FormatInt(count, 10))
		hashed := fmt.Sprintf("%x", md5.Sum(data))
		if hashed[:5] == "00000" {
			pos, err := strconv.Atoi(string(hashed[5]))
			if err == nil && pos < 8 && password[pos] == "_" {
				char := string(hashed[6])
				password[pos] = char
				passLength++
				fmt.Printf("\rDECRYPTING: %v", strings.Join(password, ""))
			}
		}
		count++
	}
	fmt.Printf("\rDONE:       %v\n", strings.Join(password, ""))
}

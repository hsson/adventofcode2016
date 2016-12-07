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
	fmt.Print("DECRYPTING:  ")
	for len(password) < 8 {
		data := []byte(input + strconv.Itoa(count))
		hashed := fmt.Sprintf("%x", md5.Sum(data))
		char := string(hashed[5])
		if count%10000 == 0 {
			// Too slow to print on every try
			fmt.Printf("\b%v", char)
		}
		if hashed[:5] == "00000" {
			password += char
			fmt.Printf("\b%v ", char)
		}
		count++
	}
	fmt.Printf("\rDONE:       %v\n", password)
}

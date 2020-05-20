package main

import (
	"fmt"
	"unicode"
)

func cipher(target string) []string {
	var ciphered []string
	for _, v := range target {
		if unicode.IsLower(v) && unicode.IsLetter(v) {
			v = 219 - v
			ciphered = append(ciphered, string(v))
			continue
		}
		fmt.Println(string(v))
	}
	return ciphered
}

func main() {
	target := "zyxwあ"
	fmt.Println(cipher(target))
}

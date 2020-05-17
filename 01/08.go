package main

import (
	"fmt"
	"unicode"
)

func cipher(target string) []string {
	var ciphered []string
	for _, v := range target {
		if unicode.IsLower(v) && unicode.IsLetter(v) {
			fmt.Println(v)
			ciphered = append(ciphered, string(v))
		}
	}
	return ciphered
}

func main() {
	target := "ほげほげ"
	fmt.Println(cipher(target))
}

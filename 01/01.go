package main

import (
	"fmt"
	"strings"
)

const word = "パタトクカシーー"

func main() {
	var chars []string
	for i, v := range word {
		if i%2 == 0 {
			continue
		}
		chars = append(chars, string(v))
	}
	ans := strings.Join(chars, "")
	fmt.Println((ans))
}

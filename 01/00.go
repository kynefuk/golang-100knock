package main

import (
	"fmt"
	"strings"
)

const word = "stressed"

func main() {
	var strArray []string
	for i := len(word) - 1; i >= 0; i-- {
		strArray = append(strArray, string(word[i]))
	}
	fmt.Println(strings.Join(strArray, ""))
}

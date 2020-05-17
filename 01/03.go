package main

import (
	"fmt"
	"strings"
)

const word = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

func main() {
	splited := strings.Split(word, " ")
	var initials []string
	for _, v := range splited {
		initials = append(initials, string(v[0]))
	}
	fmt.Println(initials)
}

package main

import (
	"fmt"
	"strings"
)

const word = `Hi He Lied Because Boron Could Not Oxidize Fluorine. 
New Nations Might Also Sign Peace Security Clause. 
Arthur King Can.`

func main() {
	splited := strings.Split(word, " ")
	splitedMap := map[int]string{}
	for i, v := range splited {
		if i%2 == 0 {
			splitedMap[i] = v[:2]
			continue
		}
		splitedMap[i] = v[:1]
	}
	fmt.Println(splitedMap)
}

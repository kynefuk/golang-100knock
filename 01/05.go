package main

import (
	"fmt"
	"strings"
)

const word = "I am an NLPer"

func Ngram(target []string, n int) []string {
	var grams []string
	for i := 0; i <= len(target)-n; i++ {
		extracted := target[i : i+n]
		grams = append(grams, strings.Join(extracted, ""))
	}
	return grams
}

func main() {
	// 単語bi-gram
	words := strings.Split(word, " ")
	fmt.Println(Ngram(words, 2))

	// 文字bi-gram
	var chars []string
	replaced := strings.ReplaceAll(word, " ", "")
	for _, v := range replaced {
		chars = append(chars, string(v))
	}

	fmt.Println(Ngram(chars, 2))
}

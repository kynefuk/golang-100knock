package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"
)

const word = "I couldn’t believe that I could actually understand what I was reading : the phenomenal power of the human mind ."

// 文字列を配列にする関数
func Split(target string) []string {
	var splited []string
	splited = strings.Split(target, " ")
	return splited
}

func MakeRandom(target []string) []string {
	length := len(target)
	randomStr := make([]string, length)
	rand.Seed(time.Now().Unix())
	randomInt := rand.Perm(length)

	for i, v := range target {
		if i == 0 || i == length-1 {
			continue
		}
		randomStr = append(randomStr, "")
		index := randomInt[i]
		copy(randomStr[index+1:], randomStr[index:])
		randomStr[index] = string(v)
	}
	fmt.Println(randomStr)
	randomStr = append(randomStr[:1], randomStr[:]...)
	randomStr[0] = target[0]
	randomStr = append(randomStr, target[length-1])
	return randomStr
}

func Typoglycemia(target string) string {
	if utf8.RuneCountInString(target) < 4 {
		return target
	}
	splited := Split(target)
	randomStr := MakeRandom(splited)
	fmt.Println(randomStr)
	str := strings.Join(randomStr, " ")
	return str
}

func main() {
	fmt.Println(Typoglycemia(word))
}

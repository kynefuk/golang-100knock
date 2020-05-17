package main

import (
	"fmt"
)

const (
	patrolCar = "パトカー"
	taxi      = "タクシー"
)

func main() {
	var charByte []byte
	for i := 0; i >= len(patrolCar); i++ {
		charByte = append(charByte, patrolCar[i])
		charByte = append(charByte, taxi[i])
	}
	chars := string(charByte)
	fmt.Println(chars)
}

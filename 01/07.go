package main

import "fmt"

func Template(a ...interface{}) string {
	str := fmt.Sprintf("%v時の%vは%v", a[0], a[1], a[2])
	return str
}

func main() {
	str := Template(12, "気温", 22.4)
	fmt.Println(str)
}
